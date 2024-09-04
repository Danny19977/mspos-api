package sup

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetSups(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.Sup{}, p, l))
}
 

// query data
func GetSupASMByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var sup []models.Sup
	db.Where("asm_id = ?", id).Find(&sup)
	 
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Sup by id found", 
		"data": sup,
	})
} 

// query data
func GetSupByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var sup []models.Sup
	db.Where("province_id = ?", id).Find(&sup)
	 
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Sup by id found", 
		"data": sup,
	})
} 
 


// Get one data
func GetSup(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var sup models.Sup
	db.Find(&sup, id)
	if sup.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No sup name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "sup found",
			"data":    sup,
		},
	)
}

// Create data
func CreateSup(c *fiber.Ctx) error {
	p := &models.Sup{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

// Update data
func UpdateSup(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name       string `json:"name"`
		ProvinceID uint   `json:"province_id"`
		AsmID      uint   `json:"asm_id"`
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

	sup := new(models.Sup)

	db.First(&sup, id)
	sup.Name = updateData.Name
	sup.ProvinceID = updateData.ProvinceID
	sup.AsmID = updateData.AsmID
	sup.Signature = updateData.Signature

	db.Save(&sup)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "sup updated success",
			"data":    sup,
		},
	)

}

// Delete data
func DeleteSup(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var sup models.Sup
	db.First(&sup, id)
	if sup.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No sup name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&sup)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "sup deleted success",
			"data":    nil,
		},
	)
}
