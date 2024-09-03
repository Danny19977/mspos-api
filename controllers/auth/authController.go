package authentification

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
	"github.com/kgermando/mspos-api/utils"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func Register(c *fiber.Ctx) error {

	nu := new(models.User)

	if err := c.BodyParser(&nu); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if nu.Password != nu.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	u := &models.User{
		Fullname:   nu.Fullname,
		Email:      nu.Email,
		Phone:      nu.Phone,
		AreaID:     nu.AreaID,
		ProvinceID: nu.ProvinceID,
		SupID:      nu.SupID,
		PosID:      nu.PosID,
		Role:       nu.Role,
		Permission: nu.Permission,
		Image:      nu.Image,
		Status:     nu.Status,
		Signature:  nu.Signature,
	}

	u.SetPassword(nu.Password)

	if err := utils.ValidateStruct(*u); err != nil {
		c.Status(400)
		return c.JSON(err)
	}

	if err := database.DB.Create(u).Error; err != nil {
		c.Status(500)
		sm := strings.Split(err.Error(), ":")
		m := strings.TrimSpace(sm[1])

		return c.JSON(fiber.Map{
			"message": m,
		})
	}

	return c.JSON(fiber.Map{
		"message": "user account created",
	})
}

func Login(c *fiber.Ctx) error {

	lu := new(models.Login)

	if err := c.BodyParser(&lu); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.ValidateStruct(*lu); err != nil {
		c.Status(400)
		return c.JSON(err)
	}

	u := &models.User{}

	database.DB.Where("email = ?", lu.Email).First(&u)

	if u.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "invalid login credentials email 😰",
		})
	}

	if err := u.ComparePassword(lu.Password); err != nil {

		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid login credentials 😰",
		})
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(u.ID)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), //1 day ,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})

}

func AuthUser(c *fiber.Ctx) error {

	cookie := c.Cookies("token")

	userId, _ := utils.VerifyJwt(cookie)

	u := models.User{}

	database.DB.Where("id = ?", userId).First(&u)

	r := &models.UserResponse{
		Id:         u.ID,
		Fullname:   u.Fullname,
		Email:      u.Email,
		Title:      u.Title,
		Phone:      u.Phone,
		Role:       u.Role,
		Area:       u.AreaID,
		Province:   u.ProvinceID,
		Sup:        u.SupID,
		Pos:        u.PosID,
		Permission: u.Permission,
		Status:     u.Status,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}

	return c.JSON(r)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //1 day ,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

// User bioprofile
func UpdateInfo(c *fiber.Ctx) error {
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
		Signature  string `json:"signature"`
	}
	var updateData UpdateDataInput

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"errors":  err.Error(),
		})
	}

	cookie := c.Cookies("token")

	Id, _ := utils.VerifyJwt(cookie)

	userId, _ := strconv.Atoi(Id)

	user := new(models.User)

	db := database.DB

	db.First(&user, userId)
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
	user.Signature = updateData.Signature

	db.Save(&user)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully updated",
		"data":    user,
	})

}

func UpdatePassword(c *fiber.Ctx) error {
	type UpdateDataInput struct {
		Password        string `json:"password"`
		PasswordConfirm string `json:"PasswordConfirm"`
	}
	var updateData UpdateDataInput

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"errors":  err.Error(),
		})
	}

	if updateData.Password != updateData.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	cookie := c.Cookies("token")

	Id, _ := utils.VerifyJwt(cookie)
	userId, _ := strconv.Atoi(Id)

	user := new(models.User)

	p, err := utils.HashPassword(updateData.Password)
	if err != nil {
		return err
	}

	db := database.DB

	db.First(&user, userId)
	user.Password = p

	// successful update remove cookies
	rmCookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //1 day ,
		HTTPOnly: true,
	}
	c.Cookie(&rmCookie)

	return c.JSON(user)

}
