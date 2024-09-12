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
		Eq        int64  `json:"eq"`
		Eq1       int64  `json:"eq1"`
		Sold      int64  `json:"sold"`
		Dhl       int64  `json:"dhl"`
		Dhl1      int64  `json:"dhl1"`
		Ar        int64  `json:"ar"`
		Ar1       int64  `json:"ar1"`
		Sbl       int64  `json:"sbl"`
		Sbl1      int64  `json:"sbl1"`
		Pmf       int64  `json:"pmf"`
		Pmf1      int64  `json:"pmf1"`
		Pmm       int64  `json:"pmm"`
		Pmm1      int64  `json:"pmm1"`
		Ticket    int64  `json:"ticket"`
		Ticket1   int64  `json:"ticket1"`
		Mtc       int64  `json:"mtc"`
		Mtc1      int64  `json:"mtc1"`
		Ws        int64  `json:"ws"`
		Ws1       int64  `json:"ws1"`
		Mast      int64  `json:"mast"`
		Mast1     int64  `json:"mast1"`
		Oris      int64  `json:"oris"`
		Oris1     int64  `json:"oris1"`
		Elite     int64  `json:"elite"`
		Elite1    int64  `json:"elite1"`
		Yes       int64  `json:"yes"`
		Yes1      int64  `json:"yes1"`
		Time      int64  `json:"time"`
		Time1     int64  `json:"time1"`
		Comment   string `json:"comment"`
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
	posform.Eq = updateData.Eq
	posform.Eq1 = updateData.Eq1
	posform.Sold = updateData.Sold
	posform.Dhl = updateData.Dhl
	posform.Dhl1 = updateData.Dhl1
	posform.Ar = updateData.Ar
	posform.Ar1 = updateData.Ar1
	posform.Sbl = updateData.Sbl
	posform.Sbl1 = updateData.Sbl1
	posform.Pmf = updateData.Pmf
	posform.Pmf1 = updateData.Pmf1
	posform.Pmm = updateData.Pmm
	posform.Pmm1 = updateData.Pmm1
	posform.Ticket = updateData.Ticket
	posform.Ticket1 = updateData.Ticket1
	posform.Mtc = updateData.Mtc
	posform.Mtc1 = updateData.Mtc1
	posform.Ws = updateData.Ws
	posform.Ws1 = updateData.Ws1
	posform.Mast = updateData.Mast
	posform.Mast1 = updateData.Mast1
	posform.Oris = updateData.Oris
	posform.Oris1 = updateData.Oris1
	posform.Elite = updateData.Elite
	posform.Elite1 = updateData.Elite1
	posform.Yes = updateData.Yes
	posform.Yes1 = updateData.Yes1
	posform.Time = updateData.Time
	posform.Time1 = updateData.Time1
	posform.Comment = updateData.Comment 
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
