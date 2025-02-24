package dr

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedDr(c *fiber.Ctx) error {
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

	var dataList []models.Dr

	var length int64
	db.Model(dataList).Count(&length)
	db.
		Where("signature ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("Dr.updated_at DESC").
		Preload("Province").
		Preload("Area").
		Preload("Asm").
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
		"message":    "All DRs",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllDr(c *fiber.Ctx) error {
	db := database.DB

	var data []models.Dr
	db.
		Preload("Province").
		Preload("Area").
		Preload("Asm").
		Preload("Sup").
		Order("Dr.updated_at DESC").
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Drs",
		"data":    data,
	})
}

// Get one data
func GetDr(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var Dr models.Dr
	db.Find(&Dr, id)
	if Dr.Signature == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Dr signature found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Dr found",
			"data":    Dr,
		},
	)
}

// Create data
func CreateDr(c *fiber.Ctx) error {
	p := &models.Dr{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Dr created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateDr(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		CountryID  uint   `json:"country_id"`
		ProvinceID uint   `json:"province_id"`
		AreaID     uint   `json:"area_id"`
		AmsID      uint   `json:"ams_id"`
		SupID      uint   `json:"sup_id"`
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

	Dr := new(models.Dr)

	db.First(&Dr, id)
	Dr.CountryID = updateData.CountryID
	Dr.ProvinceID = updateData.ProvinceID
	Dr.AreaID = updateData.AreaID
	Dr.AmsID = updateData.AmsID
	Dr.SupID = updateData.SupID
	Dr.Signature = updateData.Signature

	db.Save(&Dr)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Dr updated success",
			"data":    Dr,
		},
	)

}

// Delete data
func DeleteDr(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var Dr models.Dr
	db.First(&Dr, id)
	if Dr.Signature == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Dr name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&Dr)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Dr deleted success",
			"data":    nil,
		},
	)
}
