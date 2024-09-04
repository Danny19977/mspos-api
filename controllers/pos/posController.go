package poss

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetPoss(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.Pos{}, p, l))
}

// query data
func GetPosByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var poss []models.Pos
	db.Where("province_id = ?", id).Find(&poss)
	 
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "poss by id found", 
		"data": poss,
	})
}

// query data Area
func GetPosAreaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var poss []models.Pos
	db.Where("area_id = ?", id).Find(&poss)
	 
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "poss by id found", 
		"data": poss,
	})
}



// Get one data
func GetPos(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var pos models.Pos
	db.Find(&pos, id)
	if pos.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No pos  name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "pos  found",
			"data":    pos,
		},
	)
}

// Create data
func CreatePos(c *fiber.Ctx) error {
	p := &models.Pos{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

// Update data
func UpdatePos(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name               string `json:"name"`    // Celui qui vend
		Shop               string `json:"shop"`    // Nom du shop
		Manager            string `json:"manager"` // name of the onwer of the pos
		Commune            string `json:"commune"`
		Avenue             string `json:"avenue"`
		Quartier           string `json:"quartier"`
		Reference          string `json:"reference"`
		Telephone          int64  `json:"telephone"`
		Eparasol           string `json:"eparasol"`
		Etable             string `json:"etable"`
		Ekiosk             bool   `json:"ekiosk"`
		InputGroupSelector string `json:"inputgroupselector"`
		Cparasol           string `json:"cparasol"`
		Ctable             string `json:"ctable"`
		Ckiosk             bool   `json:"Ckiosk"`
		ProvinceID         uint   `json:"province_id"`
		AreaID             uint   `json:"area_id"`
		Status             bool   `json:"status"`
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

	pos := new(models.Pos)

	db.First(&pos, id)
	pos.Name = updateData.Name
	pos.Shop = updateData.Shop
	pos.Manager = updateData.Manager
	pos.Commune = updateData.Commune
	pos.Avenue = updateData.Avenue
	pos.Quartier = updateData.Quartier
	pos.Reference = updateData.Reference
	pos.Telephone = updateData.Telephone
	pos.Eparasol = updateData.Eparasol
	pos.Etable = updateData.Etable
	pos.Ekiosk = updateData.Ekiosk
	pos.InputGroupSelector = updateData.InputGroupSelector
	pos.Cparasol = updateData.Cparasol
	pos.Ctable = updateData.Ctable
	pos.Ckiosk = updateData.Ckiosk
	pos.ProvinceID = updateData.ProvinceID
	pos.AreaID = updateData.AreaID
	pos.Status = updateData.Status
	pos.Signature = updateData.Signature

	db.Save(&pos)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "POS updated success",
			"data":    pos,
		},
	)

}

// Delete data
func DeletePos(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var pos models.Pos
	db.First(&pos, id)
	if pos.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No POS name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&pos)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "POS deleted success",
			"data":    nil,
		},
	)
}
