package userr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
	"github.com/kgermando/mspos-api/utils"
)

// Paginate
func GetPaginatedUsers(c *fiber.Ctx) error {
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

	var dataList []models.User

	var length int64
	var data []models.User
	db.Model(data).Count(&length)

	db.
		Joins("JOIN provinces ON users.province_id=provinces.id").
		Joins("JOIN areas ON users.area_id=areas.id").
		Joins("JOIN sups ON users.sup_id=sups.id").
		// Preload("provinces").
		// Preload("areas").
		// Preload("sups").
		Where("fullname ILIKE ? OR title ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
		users.id AS id, fullname, title, email, phone, 
		provinces.name AS province, areas.name AS area, 
		sups.name AS sup, status
	`).
		Offset(offset).
		Limit(limit).
		Order("users.updated_at DESC").
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
		"message":    "All users",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data province
func GetUserByProvinceID(c *fiber.Ctx) error {
	provinceId := c.Params("province_id")

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

	var dataList []models.User

	var length int64
	var data []models.User
	db.Model(data).Where("province_id = ?", provinceId).Count(&length)

	db.
		Joins("JOIN provinces ON users.province_id=provinces.id").
		Joins("JOIN areas ON users.area_id=areas.id").
		Joins("JOIN sups ON users.sup_id=sups.id"). 
		Where("users.province_id = ?", provinceId).
		Where("fullname ILIKE ? OR title ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
		users.id AS id, fullname, title, email, phone, 
		provinces.name AS province, areas.name AS area, 
		sups.name AS sup, status
	`).
		Offset(offset).
		Limit(limit).
		Order("users.updated_at DESC").
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
		"message":    "All users",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data sup by area
func GetUserBySupID(c *fiber.Ctx) error {
	areaId := c.Params("sup_id")

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

	var dataList []models.User

	var length int64
	var data []models.User
	db.Model(data).Where("area_id = ?", areaId).Count(&length)

	db.
		Joins("JOIN provinces ON users.province_id=provinces.id").
		Joins("JOIN areas ON users.area_id=areas.id").
		Joins("JOIN sups ON users.sup_id=sups.id"). 
		Where("users.area_id = ?", areaId).
		Where("fullname ILIKE ? OR title ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
		users.id AS id, fullname, title, email, phone, 
		provinces.name AS province, areas.name AS area, 
		sups.name AS sup, status
	`).
		Offset(offset).
		Limit(limit).
		Order("users.updated_at DESC").
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
		"message":    "All users",
		"data":       dataList,
		"pagination": pagination,
	})
}


// query all data
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User
	db.Find(&users)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All users",
		"data":    users,
	})
}

// query data
func GetUserByID(c *fiber.Ctx) error {
	provinceId := c.Params("id")
	db := database.DB
	var users []models.User
	db.Where("province_id = ?", provinceId).Where("role = ?", "DR").Where("status = ?", true).Find(&users)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "users by id found",
		"data":    users,
	})
}

// Get one data
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user models.User
	db.Find(&user, id)
	if user.Fullname == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No User name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "User found",
			"data":    user,
		},
	)
}

// Create data
func CreateUser(c *fiber.Ctx) error {
	p := &models.User{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	if p.Fullname == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Form not complete",
				"data":    nil,
			},
		)
	}

	if p.Password != p.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := &models.User{
		Fullname:   p.Fullname,
		Email:      p.Email,
		Title:      p.Title,
		Phone:      p.Phone,
		AreaID:     p.AreaID,
		ProvinceID: p.ProvinceID,
		SupID:      p.SupID,
		Role:       p.Role,
		Permission: p.Permission,
		Image:      p.Image,
		Status:     p.Status,
		IsManager:  p.IsManager,
		Signature:  p.Signature,
	}

	user.SetPassword(p.Password)

	if err := utils.ValidateStruct(*user); err != nil {
		c.Status(400)
		return c.JSON(err)
	}

	if err := database.DB.Create(user).Error; err != nil {
		c.Status(500)
		sm := strings.Split(err.Error(), ":")
		m := strings.TrimSpace(sm[1])

		return c.JSON(fiber.Map{
			"message": m,
		})
	}

	// database.DB.Create(user)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "User Created success",
			"data":    user,
		},
	)
}

// Update data
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateDataInput struct {
		Fullname   string `json:"fullname"`
		Email      string `json:"email"`
		Title      string `json:"title"`
		Phone      string `json:"phone"`
		AreaID     uint   `json:"area_id"`
		ProvinceID uint   `json:"province_id"`
		SupID      uint   `json:"sup_id"`
		Role       string `json:"role"`
		Permission string `json:"permission"`
		Image      string `json:"image"`
		Status     bool   `json:"status"`
		IsManager  bool   `json:"is_manager"`
		Signature  string `json:"signature"`
	}
	var updateData UpdateDataInput

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Review your iunput",
				"data":    nil,
			},
		)
	}

	user := new(models.User)

	db.First(&user, id)
	user.Fullname = updateData.Fullname
	user.Email = updateData.Email
	user.Title = updateData.Title
	user.Phone = updateData.Phone
	user.AreaID = updateData.AreaID
	user.ProvinceID = updateData.ProvinceID
	user.SupID = updateData.SupID
	user.Role = updateData.Role
	user.Permission = updateData.Permission
	user.Image = updateData.Image
	user.Status = updateData.Status
	user.IsManager = updateData.IsManager
	user.Signature = updateData.Signature

	db.Save(&user)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "User updated success",
			"data":    user,
		},
	)

}

// Delete data
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var User models.User
	db.First(&User, id)
	if User.Fullname == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No User name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&User)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "User deleted success",
			"data":    nil,
		},
	)
}
