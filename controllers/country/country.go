package country

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	// "github.com/kgermando/mspos-api/controllers/country"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedCountry(c *fiber.Ctx) error {
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
		country.id AS id, 
		contries.name AS name 
	`).
		Offset(offset).
		Limit(limit).
		Order("countries.updated_at DESC").
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
		"message":    "All Countries",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All Country Dropdown
func GetCountryDropdown(c *fiber.Ctx) error {
	db := database.DB
	sql1 := `
		SELECT "countries"."id" AS id, "countries"."name" AS name
		FROM pos_forms
		INNER JOIN countries ON pos_forms.country_id=countries.id
		WHERE "pos_forms"."deleted_at" IS NULL
		GROUP BY "countries"."id", "countries"."name"; 
	`
	var data []models.ProvinceDropDown
	db.Raw(sql1).Scan(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All countries Dropdown",
		"data":    data,
	})
}

// Get All data
func GetAllCountry(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Country
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All countries",
		"data":    data,
	})
}

// query data
func GetCountryeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var countries []models.Country
	db.Where("id = ?", id).Find(&countries)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "countries by id found",
		"data":    countries,
	})
}

// Get one data
func GetCountry(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var Country models.Country
	db.Find(&Country, id)
	if Country.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Country name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Country found",
			"data":    Country,
		},
	)
}

// Create data
func CreateCountry(c *fiber.Ctx) error {
	p := &models.Country{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Country created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateCountry(c *fiber.Ctx) error {
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

	country := new(models.Country)

	db.First(&country, id)
	country.Name = updateData.Name
	country.Signature = updateData.Signature

	db.Save(&country)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Country updated success",
			"data":    country,
		},
	)

}

// Delete data
func DeleteCountry(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var country models.Country
	db.First(&country, id)
	if country.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Country name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&country)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Country deleted success",
			"data":    nil,
		},
	)
}
