package asm

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetAsms(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.Asm{}, p, l))
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

	return c.JSON(p)
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
