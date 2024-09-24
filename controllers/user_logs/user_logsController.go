package userlogs

import ( 
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetUserLogs(c *fiber.Ctx) error {

	db := database.DB
	var data []models.UserLogs
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All UserLogs",
		"data":    data,
	})
}


// Get one data
func GetUserLog(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user_logs models.UserLogs
	db.Find(&user_logs, id)
	if user_logs.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No user_logs  name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "user_logs  found",
			"data":    user_logs,
		},
	)
}

// Create data
func CreateUserLog(c *fiber.Ctx) error {
	p := &models.UserLogs{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

// Update data
func UpdateUserLog(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name        string `json:"name"`
		UserID      uint `json:"user_id"`
		Action      string `json:"action"`
		Description string `json:"description"`
		Signature    string `json:"signature"`
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

	user_logs := new(models.UserLogs)

	db.First(&user_logs, id)
	user_logs.Name = updateData.Name
	user_logs.UserID = updateData.UserID
	user_logs.Action = updateData.Action
	user_logs.Description = updateData.Description
	user_logs.Signature = updateData.Signature

	db.Save(&user_logs)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "user_logs  updated success",
			"data":    user_logs,
		},
	)

}

// Delete data
func DeleteUserLog(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var user_logs models.UserLogs
	db.First(&user_logs, id)
	if user_logs.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No user_logs  name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&user_logs)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "user_logs  deleted success",
			"data":    nil,
		},
	)
}
