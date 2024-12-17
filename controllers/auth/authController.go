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
		Title:      nu.Title,
		Phone:      nu.Phone,
		AreaID:     nu.AreaID,
		ProvinceID: nu.ProvinceID,
		SupID:      nu.SupID,
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
			"message": "invalid email 😰",
		})
	}

	if err := u.ComparePassword(lu.Password); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "mot de passe incorrect! 😰",
		})
	}

	if !u.Status {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "vous n'êtes pas autorisé de se connecter 😰",
		})
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(u.ID)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:   "token",
		Value:  token,
		Path:   "/",
		// Domain: ".render.com",
		// Domain: "localhost",
		Expires:  time.Now().Add(time.Hour * 24), //1 day ,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "none",
		// SessionOnly: true,
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
		Expires:  time.Now().Add(-time.Hour), // 1 day ,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"Logout":  "success",
	})

}

// User bioprofile
func UpdateInfo(c *fiber.Ctx) error {
	type UpdateDataInput struct {
		Fullname  string `json:"fullname"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Signature string `json:"signature"`
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
	user.Phone = updateData.Phone
	user.Signature = updateData.Signature

	db.Save(&user)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully updated",
		"data":    user,
	})

}

func ChangePassword(c *fiber.Ctx) error {
	type UpdateDataInput struct {
		OldPassword     string `json:"old_password"`
		Password        string `json:"password"`
		PasswordConfirm string `json:"password_confirm"`
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

	userId, _ := utils.VerifyJwt(cookie)

	user := new(models.User)

	database.DB.Where("id = ?", userId).First(&user)

	if err := user.ComparePassword(updateData.OldPassword); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "votre mot de passe n'est pas correct! 😰",
		})
	}

	if updateData.Password != updateData.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	p, err := utils.HashPassword(updateData.Password)
	if err != nil {
		return err
	}

	db := database.DB

	db.First(&user, user.ID)
	user.Password = p

	db.Save(&user)

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
