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

	var dataList []models.Sup

	var length int64
	// var data []models.Sup
	db.Model(dataList).Count(&length)

	db.
		Joins("JOIN provinces ON sups.province_id=provinces.id").
		Joins("JOIN asms ON sups.asm_id=asms.id").
		Where("asms.name ILIKE ?", "%"+search+"%").
		Select(`
		sups.id AS id, 
		sups.name AS name, 
		provinces.name AS province,
		asms.name AS asm
	`).
		Offset(offset).
		Limit(limit).
		Order("sups.updated_at DESC").
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
		"message":    "All sups",
		"data":       dataList,
		"pagination": pagination,
	})
}

// query data province
func GetSupByProvinceID(c *fiber.Ctx) error {
	db := database.DB
	provinceId := c.Params("province_id")

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

	var dataList []models.Sup

	var length int64
	var data []models.Sup
	db.Model(data).Where("province_id = ?", provinceId).Count(&length) 

	db.
		Joins("JOIN provinces ON sups.province_id=provinces.id").
		Joins("JOIN asms ON sups.asm_id=asms.id").
		Where("sups.province_id = ?", provinceId).
		Where("asms.name ILIKE ?", "%"+search+"%").
		Select(`
		sups.id AS id, 
		sups.name AS name, 
		provinces.name AS province,
		asms.name AS asm
	`).
		Offset(offset).
		Limit(limit).
		Order("sups.updated_at DESC").
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
		"message":    "All sup by province",
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
	asm_id := c.Params("asm_id")
	db := database.DB
	var sup []models.Sup
	db.Where("asm_id = ?", asm_id).Find(&sup)
	 
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

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Sup created success",
			"data":    p,
		},
	)
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
