package Subarea

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
	"github.com/lib/pq"
)

// Paginate
func GetPaginatedSubArea(c *fiber.Ctx) error {
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

	var dataList []models.SubArea

	var length int64
	db.Model(dataList).Count(&length)
	db.Where("name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("Subarea.updated_at DESC").
		Preload("Province").
		Preload("Sup").
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
		"message":    "All products",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllSubArea(c *fiber.Ctx) error {
	db := database.DB

	var data []models.SubArea
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Subarea",
		"data":    data,
	})
}

// Get All data by id
func GetAllSubAreaBySearch(c *fiber.Ctx) error {
	db := database.DB

	search := c.Query("search", "")

	var data []models.SubArea
	db.Where("name ILIKE ?", "%"+search+"%").
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Subarea",
		"data":    data,
	})
}

// Get one data
func GetSubArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var SubArea models.SubArea
	db.Find(&SubArea, id)
	if SubArea.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Subarea name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Subarea found",
			"data":    SubArea,
		},
	)
}

// Create data
func CreateSubArea(c *fiber.Ctx) error {
	p := &models.SubArea{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Subarea created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateSubArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name       string         `json:"name"`
		Commune    pq.StringArray `json:"commune"`
		ProvinceID uint           `json:"province_id"`
		SupID      uint           `json:"sup_id"`
		Signature  string         `json:"signature"`
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

	SubArea := new(models.SubArea)

	db.First(&SubArea, id)
	SubArea.Name = updateData.Name

	db.Save(&SubArea)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Subarea updated success",
			"data":    SubArea,
		},
	)

}

// Delete data
func DeleteSubarea(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var SubArea models.SubArea
	db.First(&SubArea, id)
	if SubArea.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Subarea name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&SubArea)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Subarea deleted success",
			"data":    nil,
		},
	)
}
