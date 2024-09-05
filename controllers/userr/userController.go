package userr

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
	"github.com/kgermando/mspos-api/utils"
)

// Get All data
func GetUsers(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, p, l))
}


// query data
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var users []models.User
	db.Where("province_id = ?", id).Find(&users)
	 
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "users by id found", 
		"data": users,
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
		PosID:      p.PosID,
		Role:       p.Role,
		Permission: p.Permission,
		Image:      p.Image,
		Status:     p.Status,
		IsManager: p.IsManager,
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
		PosID      uint   `json:"pos_id"`
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
	user.PosID = updateData.PosID
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
