package area

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetAreas(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.Area{}, p, l))
}

// Get one data
func GetArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var area models.Area
	db.Find(&area, id)
	if area.Shop == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No area name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "area found",
			"data":    area,
		},
	)
}

// Create data
func CreateArea(c *fiber.Ctx) error {
	p := &models.Area{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

// Update data
func UpdateArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Shop               string `json:"shop"`
		Manager            string `json:"manager"` // name of the onwer of the pos
		Commune            string `json:"commune"`
		Avenue             string `json:"avenue"`
		Quartier           string `json:"quartier"`
		Reference          string `json:"reference"`
		Number             int64  `json:"number"`
		Eparasol           string `json:"eparasol"`
		Etable             string `json:"etable"`
		Ekiosk             bool   `json:"ekiosk"`
		InputGroupSelector string `json:"inputgroupselector"`
		Cparasol           string `json:"cparasol"`
		Ctable             string `json:"ctable"`
		Ckiosk             bool   `json:"Ckiosk"`
		ProvinceID         uint   `json:"province_id"`
		UserID             uint   `json:"user_id"`
		Signature          string `json:"signature"`
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

	area := new(models.Area)

	db.First(&area, id)
	area.Shop = updateData.Shop
	area.Manager = updateData.Manager
	area.Commune = updateData.Commune
	area.Avenue = updateData.Avenue
	area.Quartier = updateData.Quartier
	area.Reference = updateData.Reference
	area.Number = updateData.Number
	area.Eparasol = updateData.Eparasol
	area.Etable = updateData.Etable
	area.Ekiosk = updateData.Ekiosk
	area.InputGroupSelector = updateData.InputGroupSelector
	area.Cparasol = updateData.Cparasol
	area.Ctable = updateData.Ctable
	area.Ckiosk = updateData.Ckiosk
	area.ProvinceID = updateData.ProvinceID
	area.UserID = updateData.UserID
	area.Signature = updateData.Signature

	db.Save(&area)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "area updated success",
			"data":    area,
		},
	)

}

// Delete data
func DeleteArea(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var area models.Area
	db.First(&area, id)
	if area.Shop == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No area name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&area)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "area deleted success",
			"data":    nil,
		},
	)
}
