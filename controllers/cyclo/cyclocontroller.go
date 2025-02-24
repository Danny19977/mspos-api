package cyclo

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedCyclo(c *fiber.Ctx) error {
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

	var dataList []models.Cyclo

	var length int64
	db.Model(dataList).Count(&length)
	db.
		Where("name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("cyclo.updated_at DESC").
		Preload("Province").
		Preload("Area").
		Preload("Asm").
		Preload("Sup").
		Preload("Dr").
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
func GetAllCyclo(c *fiber.Ctx) error {
	db := database.DB

	var data []models.Cyclo
	db.
		Preload("Province").
		Preload("Area").
		Preload("Asm").
		Preload("Sup").
		Preload("Dr").
		Order("Cyclo.updated_at DESC").
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Cyclos",
		"data":    data,
	})
}

// Get one data
func GetCyclo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var cyclo models.Cyclo
	db.Find(&cyclo, id)
	if cyclo.Signature == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No cyclo signature found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "cyclo found",
			"data":    cyclo,
		},
	)
}

// Create data
func CreateCyclo(c *fiber.Ctx) error {
	p := &models.Cyclo{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Cyclo created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateCyclo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		CountryID  uint   `json:"country_id"`
		ProvinceID uint   `json:"province_id"`
		AreaID     uint   `json:"area_id"`
		AmsID      uint   `json:"ams_id"`
		SupID      uint   `json:"sup_id"`
		DrID       uint   `json:"dr_id"`
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

	cyclo := new(models.Cyclo)

	db.First(&cyclo, id)
	cyclo.CountryID = updateData.CountryID
	cyclo.ProvinceID = updateData.ProvinceID
	cyclo.AreaID = updateData.AreaID
	cyclo.AmsID = updateData.AmsID
	cyclo.SupID = updateData.SupID
	cyclo.DrID = updateData.DrID
	cyclo.Signature = updateData.Signature

	db.Save(&cyclo)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "cyclo updated success",
			"data":    cyclo,
		},
	)

}

// Delete data
func DeleteCyclo(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var cyclo models.Cyclo
	db.First(&cyclo, id)
	if cyclo.Signature == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Cyclo name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&cyclo)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "cyclo deleted success",
			"data":    nil,
		},
	)
}
