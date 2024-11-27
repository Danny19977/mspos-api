package manager

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedManager(c *fiber.Ctx) error {
	db := database.DB
 
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit

	search := c.Query("search", "")

	var dataList []models.Manager

	var length int64
	// var data []models.Manager
	db.Model(dataList).Count(&length) 
	db. 
		Where("name ILIKE ?", "%"+search+"%").
		Select(`
		managers.id AS id, 
		managers.name AS name 
	`).
		Offset(offset).
		Limit(limit).
		Order("managers.updated_at DESC").
		Find(&dataList)

	if err != nil {
		fmt.Println("error s'est produite: ", err)
		return c.Status(500).SendString(err.Error())
	}

	// Calculate total number of pages
	totalPages := len(dataList) / limit
	if remainder := len(dataList) % limit; remainder > 0 {
		totalPages++
	}
	pagination := map[string]interface{}{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   limit,
		"length":      length,
	} 
 
	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All PosForms",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllManagers(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Manager
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Managers",
		"data":    data,
	})
}

// Get one data
func GetManager(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var manager models.Manager
	db.Find(&manager, id)
	if manager.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No manager  name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "manager  found",
			"data":    manager,
		},
	)
}

// Create data
func CreateManager(c *fiber.Ctx) error {
	p := &models.Manager{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "manager created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateManager(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name      string `json:"name"`
		Signature string `json:"signature"`
	}

	var updateData UpdateData

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Review your iunput",
				"data":    nil,
			},
		)
	}

	manager := new(models.Manager)

	db.First(&manager, id)
	manager.Name = updateData.Name
	manager.Signature = updateData.Signature

	db.Save(&manager)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "manager  updated success",
			"data":    manager,
		},
	)

}

// Delete data
func DeleteManager(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var manager models.Manager
	db.First(&manager, id)
	if manager.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Manager name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&manager)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Manager deleted success",
			"data":    nil,
		},
	)
}
