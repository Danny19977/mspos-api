package brand

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedBrands(c *fiber.Ctx) error {
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

	var dataList []models.Brand

	var length int64
	db.Model(dataList).Count(&length)
	db.
		Where("name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("brands.updated_at DESC").
		Preload("Province").
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
func GetAllBrands(c *fiber.Ctx) error {
	db := database.DB
	
	var data []models.Brand
	db.
		Preload("Province").
		Order("brands.updated_at DESC").
	Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Brands",
		"data":    data,
	})
}

// Get one data
func GetBrand(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var brand models.Brand
	db.Find(&brand, id)
	if brand.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No brand name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "brand found",
			"data":    brand,
		},
	)
}

// Create data
func CreateBrand(c *fiber.Ctx) error {
	p := &models.Brand{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "brand created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateBrand(c *fiber.Ctx) error {
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

	brand := new(models.Brand)

	db.First(&brand, id)
	brand.Name = updateData.Name
	brand.ProvinceID = updateData.ProvinceID
	brand.Signature = updateData.Signature

	db.Save(&brand)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "brand updated success",
			"data":    brand,
		},
	)

}

// Delete data
func DeleteBrand(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var brand models.Brand
	db.First(&brand, id)
	if brand.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No product name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&brand)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "brand deleted success",
			"data":    nil,
		},
	)
}
