package manager

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetManagers(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.Manager{}, p, l))
}

// Get one data
func GetManager(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var manager models.Manager
	db.Find(&manager, id)
	if manager.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No manager  name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "manager  found",
			"data":    manager,
		},
	)
}

// Create data
func Createmanager(c *fiber.Ctx) error {
	p := &models.Manager{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

// Update data
func UpdateManager(c *fiber.Ctx) error {
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

	manager := new(models.Manager)

	db.First(&manager, id)
	manager.Name = updateData.Name
	manager.Signature = updateData.Signature

	db.Save(&manager)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "manager  updated success",
			"data":    manager,
		},
	)

}

// Delete data
func DeleteManager(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var manager models.Manager
	db.First(&manager, id)
	if manager.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Province name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&manager)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "Province deleted success",
			"data":    nil,
		},
	)
}
