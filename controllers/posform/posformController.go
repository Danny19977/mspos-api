package posform

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetPosforms(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.PosForm{}, p, l))
}

// Get one data
func GetPosform(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var posform models.PosForm
	db.Find(&posform, id)
	if posform.UserID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No posform name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "posform found",
			"data":    posform,
		},
	)
}

// Create data
func CreatePosform(c *fiber.Ctx) error {
	p := &models.PosForm{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

// Update data
func UpdatePosform(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		// IdUnique    string  `json:"id_unique"`
		Equateur int64  `json:"equateur"`
		Sold     int64  `json:"sold"`
		Dhl      int64  `json:"dhl"`
		Ar       int64  `json:"ar"`
		Sbl      int64  `json:"sbl"`
		Pmf      int64  `json:"pmf"`
		Pmm      int64  `json:"pmm"`
		Ticket   int64  `json:"ticket"`
		Mtc      int64  `json:"mtc"`
		Ws       int64  `json:"ws"`
		Mast     int64  `json:"mast"`
		Oris     int64  `json:"oris"`
		Elite    int64  `json:"elite"`
		Ck       int64  `json:"ck"`
		Yes      int64  `json:"yes"`
		Time     int64  `json:"time"`
		Comment  string `json:"comment"`
		// ProvinceID uint   `json:"province_id"`
		// AreaID     uint   `json:"area_id"`
		// SupID      uint   `json:"sup_id"`
		// PosID      uint   `json:"pos_id"`
		// UserID     uint   `json:"user_id"`
		Signature string `json:"signature"`
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

	posform := new(models.PosForm)

	db.First(&posform, id)
	// posform.IdUnique = updateData.IdUnique
	posform.Equateur = updateData.Equateur
	posform.Sold = updateData.Sold
	posform.Dhl = updateData.Dhl
	posform.Ar = updateData.Ar
	posform.Sbl = updateData.Sbl
	posform.Pmf = updateData.Pmf
	posform.Pmm = updateData.Pmm
	posform.Ticket = updateData.Ticket
	posform.Mtc = updateData.Mtc
	posform.Ws = updateData.Ws
	posform.Mast = updateData.Mast
	posform.Oris = updateData.Oris
	posform.Elite = updateData.Elite 
	posform.Yes = updateData.Yes
	posform.Time = updateData.Time
	posform.Comment = updateData.Comment
	// posform.ProvinceID = updateData.ProvinceID
	// posform.AreaID = updateData.AreaID
	// posform.SupID = updateData.SupID
	// posform.PosID = updateData.PosID
	// posform.UserID = updateData.UserID
	posform.Signature = updateData.Signature

	db.Save(&posform)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "posform updated success",
			"data":    posform,
		},
	)

}

// Delete data
func DeletePosform(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var posform models.PosForm
	db.First(&posform, id)
	if posform.UserID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No posform name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&posform)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "posform deleted success",
			"data":    nil,
		},
	)
}
