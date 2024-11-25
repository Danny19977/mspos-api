package province

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedProvince(c *fiber.Ctx) error {
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

	var dataList []models.Province

	var length int64
	// var data []models.Province
	db.Model(dataList).Count(&length)

	db. 
		Where("name ILIKE ?", "%"+search+"%").
		Select(`
		provinces.id AS id, 
		provinces.name AS name 
	`).
		Offset(offset).
		Limit(limit).
		Order("provinces.updated_at DESC").
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
		"message":    "All provinces",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All Provinces Dropdown
func GetProvinceDropdown(c *fiber.Ctx) error {
	db := database.DB
	sql1 := `
		SELECT "provinces"."id" AS id, "provinces"."name" AS name
		FROM pos_forms
		INNER JOIN provinces ON pos_forms.province_id=provinces.id
		WHERE "pos_forms"."deleted_at" IS NULL
		GROUP BY "provinces"."id", "provinces"."name"; 
	`
	var data []models.ProvinceDropDown
	db.Raw(sql1).Scan(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All provinces Dropdown",
		"data":    data,
	})
}

// Get All data
func GetAllProvinces(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Province
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All provinces",
		"data":    data,
	})
}

// query data
func GetProvinceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var provinces []models.Province
	db.Where("id = ?", id).Find(&provinces)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "provinces by id found",
		"data":    provinces,
	})
}

// Get one data
func GetProvince(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var province models.Province
	db.Find(&province, id)
	if province.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Province name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Province found",
			"data":    province,
		},
	)
}

// Create data
func CreateProvince(c *fiber.Ctx) error {
	p := &models.Province{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Province created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateProvince(c *fiber.Ctx) error {
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

	province := new(models.Province)

	db.First(&province, id)
	province.Name = updateData.Name
	province.Signature = updateData.Signature

	db.Save(&province)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Province updated success",
			"data":    province,
		},
	)

}

// Delete data
func DeleteProvince(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var province models.Province
	db.First(&province, id)
	if province.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Province name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&province)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Province deleted success",
			"data":    nil,
		},
	)
}
