package manager

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedManager(c *fiber.Ctx) error {
	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page") // CurrentPage

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 15
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	offset := (page - 1) * pageSize

	var u []models.Manager
	var length int64
	db := database.DB
	db.Find(&u).Count(&length)

	sql1 := `
		SELECT "managers"."id" AS id,  
		"managers"."name" AS name
		FROM managers   
		WHERE "managers"."deleted_at" IS NULL
		ORDER BY "managers"."updated_at" DESC;
	`
	var dataList []models.ManagerPaginate
	database.DB.Raw(sql1).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.ManagerPaginate{} // Empty slice
	} else {
		end := offset + pageSize
		if end > len(dataList) {
			end = len(dataList)
		}
		dataList = dataList[offset:end]
	}
	// Calculate total number of pages
	totalPages := len(dataList) / pageSize
	if remainder := len(dataList) % pageSize; remainder > 0 {
		totalPages++
	}

	// Create pagination metadata (adjust fields as needed)
	pagination := map[string]interface{}{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   pageSize,
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
func Createmanager(c *fiber.Ctx) error {
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
				"message": "No Province name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&manager)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Province deleted success",
			"data":    nil,
		},
	)
}
