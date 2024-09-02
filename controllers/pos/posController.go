package poss

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetPoss(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.Pos{}, p, l))
}

// Get one data
func GetPos(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var pos models.Pos
	db.Find(&pos, id)
	if pos.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No pos  name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "pos  found",
			"data":    pos,
		},
	)
}

// Create data
func CreatePos(c *fiber.Ctx) error {
	p := &models.Pos{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

// Update data
func UpdatePos(c *fiber.Ctx) error {
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

	pos := new(models.Pos)

	db.First(&pos, id)
	pos.Name = updateData.Name
	pos.Signature = updateData.Signature

	db.Save(&pos)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "manager  updated success",
			"data":    pos,
		},
	)

}

// Delete data
func DeletePos(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var pos models.Pos
	db.First(&pos, id)
	if pos.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Province name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&pos)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Province deleted success",
			"data":    nil,
		},
	)
}
