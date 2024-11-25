package userlogs

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedUserLogs(c *fiber.Ctx) error {
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

	var dataList []models.UserLogs

	var length int64
	var data []models.UserLogs
	db.Model(data).Count(&length)

	db.
		Joins("JOIN users ON user_logs.user_id=users.id").
		Where("users.fullname ILIKE ? OR user_logs.name ILIKE ? OR users.title ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
		user_logs.id AS id,  
		user_logs.name AS name, 
		user_logs.action AS action,
		user_logs.description AS description,
		user_logs.created_at AS created_at,
		users.id AS user_id,
		users.fullname AS fullname,
		users.title AS title
	`).
		Offset(offset).
		Limit(limit).
		Order("user_logs.updated_at DESC").
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
		"message":    "All UserLogs",
		"data":       dataList,
		"pagination": pagination,
	})
}

// query data
func GetUserLogByID(c *fiber.Ctx) error {
	db := database.DB
	userId := c.Params("user_id")

	
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

	var dataList []models.UserLogs

	var length int64
	var data []models.UserLogs
	db.Model(data).Where("user_id = ?", userId).Count(&length) 

	db.
		Joins("JOIN users ON user_logs.user_id=users.id").
		Where("user_logs.user_id = ?", userId).
		Where("fullname ILIKE ? OR name ILIKE ? OR title ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
		user_logs.id AS id,  
		user_logs.name AS name, 
		user_logs.action AS action,
		user_logs.description AS description,
		user_logs.created_at AS created_at,
		users.id AS user_id,
		users.fullname AS fullname,
		users.title AS title
	`).
		Offset(offset).
		Limit(limit).
		Order("user_logs.updated_at DESC").
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
		"message":    "All UserLogs",
		"data":       dataList,
		"pagination": pagination,
	})
}

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

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "UserLog created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateUserLog(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name        string `json:"name"`
		UserID      uint   `json:"user_id"`
		Action      string `json:"action"`
		Description string `json:"description"`
		Signature   string `json:"signature"`
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
