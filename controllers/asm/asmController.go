package asm

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedASM(c *fiber.Ctx) error {
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

	var u []models.Asm
	var length int64
	db := database.DB
	db.Find(&u).Count(&length)

	sql1 := `
		SELECT "asms"."id" AS id,  
		"asms"."name" AS name,
		"provinces"."name" AS province 
		FROM asms    
		INNER JOIN provinces ON asms.province_id=provinces.id 
		WHERE "asms"."deleted_at" IS NULL
		ORDER BY "asms"."updated_at" DESC;
	`
	var dataList []models.AsmPaginate
	database.DB.Raw(sql1).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.AsmPaginate{} // Empty slice
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
func GetAllAsms(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Asm
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Asms",
		"data":    data,
	})
}

// Get one data
func GetAsm(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var asm models.Asm
	db.Find(&asm, id)
	if asm.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No asm name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "asm found",
			"data":    asm,
		},
	)
}

// Create data
func CreateAsm(c *fiber.Ctx) error {
	p := &models.Asm{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "asm created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateAsm(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name       string `json:"name"`
		ProvinceID uint   `json:"province_id"`
		Signature  string `json:"signature"`
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

	asm := new(models.Asm)

	db.First(&asm, id)
	asm.Name = updateData.Name
	asm.ProvinceID = updateData.ProvinceID
	asm.Signature = updateData.Signature

	db.Save(&asm)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "asm updated success",
			"data":    asm,
		},
	)

}

// Delete data
func DeleteAsm(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var asm models.Asm
	db.First(&asm, id)
	if asm.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No asm name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&asm)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "asm deleted success",
			"data":    nil,
		},
	)
}
