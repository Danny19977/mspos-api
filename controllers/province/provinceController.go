package province

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedProvince(c *fiber.Ctx) error {
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

	var u []models.Province 
	var length int64
	db := database.DB
	db.Find(&u).Count(&length)  

	sql1 := `
		SELECT "provinces"."id" AS id, "provinces"."name" AS name 
		FROM provinces  
		ORDER BY "provinces"."updated_at" DESC;
	`
	var dataList []models.ProvincePaginate
	database.DB.Raw(sql1).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.ProvincePaginate{} // Empty slice
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
		"message":    "All provinces",
		"data":       dataList,
		"pagination": pagination,
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
		"status": "success", 
		"message": "provinces by id found", 
		"data": provinces,
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

	return c.JSON(p)
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
