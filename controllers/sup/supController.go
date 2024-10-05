package sup

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedSups(c *fiber.Ctx) error {
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

	var u []models.Sup 
	var length int64
	db := database.DB
	db.Find(&u).Count(&length) 
	fmt.Println("length", length)

	sql1 := `
		SELECT "sups"."id" AS id, "sups"."name" AS name, "provinces"."name" AS province, "asms"."name" AS asm
		FROM sups 
			INNER JOIN provinces ON sups.province_id=provinces.id 
			INNER JOIN asms ON sups.asm_id=asms.id 
			WHERE "sups"."deleted_at" IS NULL
		ORDER BY "sups"."updated_at" DESC;
	`
	var dataList []models.SupPaginate
	database.DB.Raw(sql1).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.SupPaginate{} // Empty slice
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
		"message":    "All sups",
		"data":       dataList,
		"pagination": pagination,
	})
}


// Get All data
func GetAllSups(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Sup
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All sups",
		"data":    data,
	})
}
 

// query data
func GetSupASMByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var sup []models.Sup
	db.Where("asm_id = ?", id).Find(&sup)
	 
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Sup by id found", 
		"data": sup,
	})
} 

// query data
func GetSupByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var sup []models.Sup
	db.Where("province_id = ?", id).Find(&sup)
	 
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Sup by id found", 
		"data": sup,
	})
} 
 


// Get one data
func GetSup(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var sup models.Sup
	db.Find(&sup, id)
	if sup.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No sup name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "sup found",
			"data":    sup,
		},
	)
}

// Create data
func CreateSup(c *fiber.Ctx) error {
	p := &models.Sup{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

// Update data
func UpdateSup(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name       string `json:"name"`
		ProvinceID uint   `json:"province_id"`
		AsmID      uint   `json:"asm_id"`
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

	sup := new(models.Sup)

	db.First(&sup, id)
	sup.Name = updateData.Name
	sup.ProvinceID = updateData.ProvinceID
	sup.AsmID = updateData.AsmID
	sup.Signature = updateData.Signature

	db.Save(&sup)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "sup updated success",
			"data":    sup,
		},
	)

}

// Delete data
func DeleteSup(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var sup models.Sup
	db.First(&sup, id)
	if sup.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No sup name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&sup)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "sup deleted success",
			"data":    nil,
		},
	)
}
