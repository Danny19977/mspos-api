package userlogs

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedUserLogs(c *fiber.Ctx) error {
	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page") // CurrentPage

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 15
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	offset := (page - 1) * pageSize

	var u []models.UserLogs
	var length int64
	db := database.DB
	db.Find(&u).Count(&length)

	sql1 := `
		SELECT "user_logs"."id" AS id,  
		"user_logs"."name" AS name, 
		"user_logs"."action" AS action,
		"user_logs"."description" AS description,
		"user_logs"."created_at" AS created_at,
		"users"."id" AS user_id,
		"users"."fullname" AS fullname,
		"users"."title" AS title
		FROM user_logs 
			INNER JOIN users ON user_logs.user_id=users.id   
			WHERE "user_logs"."deleted_at" IS NULL
		ORDER BY "user_logs"."updated_at" DESC;
	`
	var dataList []models.UserLogPaginate
	database.DB.Raw(sql1).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.UserLogPaginate{} // Empty slice
	} else {
		end := offset + pageSize
		if end > len(dataList) {
			end = len(dataList)
		}
		dataList = dataList[offset:end]
	}
	// Calculate total number of pages
	totalPages := len(dataList) / pageSize
	if remainder := len(dataList) % pageSize; remainder > 0 {
		totalPages++
	}

	// Create pagination metadata (adjust fields as needed)
	pagination := map[string]interface{}{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   pageSize,
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

// query data
func GetUserLogByID(c *fiber.Ctx) error {
	userId := c.Params("id")

	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page") // CurrentPage

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 15
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	offset := (page - 1) * pageSize

	var u []models.UserLogs
	var length int64
	db := database.DB
	db.Where("user_id = ?", userId).Find(&u).Count(&length)

	sql1 := `
		SELECT "user_logs"."id" AS id,  
		"user_logs"."name" AS name, 
		"user_logs"."action" AS action,
		"user_logs"."description" AS description,
		"users"."id" AS user_id,
		"users"."fullname" AS fullname,
		"users"."title" AS title
		FROM user_logs 
			INNER JOIN users ON user_logs.user_id=users.id   
			WHERE "user_logs"."deleted_at" IS NULL AND "user_logs"."user_id"=?
			ORDER BY "user_logs"."updated_at" DESC;
	`
	var dataList []models.UserLogPaginate
	database.DB.Raw(sql1, userId).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.UserLogPaginate{} // Empty slice
	} else {
		end := offset + pageSize
		if end > len(dataList) {
			end = len(dataList)
		}
		dataList = dataList[offset:end]
	}
	// Calculate total number of pages
	totalPages := len(dataList) / pageSize
	if remainder := len(dataList) % pageSize; remainder > 0 {
		totalPages++
	}

	// Create pagination metadata (adjust fields as needed)
	pagination := map[string]interface{}{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   pageSize,
		"length":      length,
	}

	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All UserLogs",
		"data":       dataList,
		"pagination": pagination,
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
