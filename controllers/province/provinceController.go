package province

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetProvinces(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.Province{}, p, l))
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
