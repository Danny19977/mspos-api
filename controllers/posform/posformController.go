package posform

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedPosForm(c *fiber.Ctx) error {
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

	var dataList []models.PosForm

	var length int64
	var data []models.PosForm
	db.Model(data).Count(&length)

	db.
		Joins("JOIN provinces ON pos_forms.province_id=provinces.id").
		Joins("JOIN sups ON pos_forms.sup_id=sups.id").
		Joins("JOIN users ON pos_forms.user_id=users.id").
		Joins("JOIN areas ON pos_forms.area_id=areas.id").
		Joins("JOIN pos ON pos_forms.pos_id=pos.id").
		Where("users.fullname ILIKE ? OR provinces.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
		pos_forms.id AS id, 
		id_unique AS id_unique, 
		eq AS eq, 
		sold AS sold, 
		dhl AS dhl, 
		ar AS ar, 
		sbl AS sbl, 
		pmf AS pmf, 
		pmm AS pmm, 
		ticket AS ticket, 
		mtc AS mtc, 
		ws AS ws, 
		mast AS mast, 
		oris AS oris, 
		elite AS elite, 
		yes AS yes, 
		time AS time, 
		comment AS comment,  
		provinces.name AS province, 
		sups.name AS sup, 
		users.fullname AS user, 
		areas.name AS area,
		pos.shop AS pos
	`).
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
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
		"message":    "All PosForms",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data DR
func GetPosformByID(c *fiber.Ctx) error {
	db := database.DB
	userId := c.Params("user_id")

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

	var dataList []models.PosForm

	var length int64
	var data []models.PosForm
	db.Model(data).Where("user_id = ?", userId).Count(&length)

	db.
		Joins("JOIN provinces ON pos_forms.province_id=provinces.id").
		Joins("JOIN sups ON pos_forms.sup_id=sups.id").
		Joins("JOIN users ON pos_forms.user_id=users.id").
		Joins("JOIN areas ON pos_forms.area_id=areas.id").
		Joins("JOIN pos ON pos_forms.pos_id=pos.id").
		Where("pos_forms.user_id = ?", userId).
		Where("users.fullname ILIKE ? OR provinces.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
			pos_forms.id AS id, 
			id_unique AS id_unique, 
			eq AS eq, 
			sold AS sold, 
			dhl AS dhl, 
			ar AS ar, 
			sbl AS sbl, 
			pmf AS pmf, 
			pmm AS pmm, 
			ticket AS ticket, 
			mtc AS mtc, 
			ws AS ws, 
			mast AS mast, 
			oris AS oris, 
			elite AS elite, 
			yes AS yes, 
			time AS time, 
			comment AS comment,  
			provinces.name AS province, 
			sups.name AS sup, 
			users.fullname AS user, 
			areas.name AS area,
			pos.shop AS pos
		`).
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
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
		"message":    "All posform by dr",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query data province
func GetPosformByProvinceID(c *fiber.Ctx) error {
	db := database.DB
	provinceId := c.Params("province_id")

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

	var dataList []models.PosForm

	var length int64
	var data []models.PosForm
	db.Model(data).Where("province_id = ?", provinceId).Count(&length)

	db.
		Joins("JOIN provinces ON pos_forms.province_id=provinces.id").
		Joins("JOIN sups ON pos_forms.sup_id=sups.id").
		Joins("JOIN users ON pos_forms.user_id=users.id").
		Joins("JOIN areas ON pos_forms.area_id=areas.id").
		Joins("JOIN pos ON pos_forms.pos_id=pos.id").
		Where("pos_forms.province_id = ?", provinceId).
		Where("users.fullname ILIKE ? OR provinces.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
			pos_forms.id AS id,
			id_unique AS id_unique, 
			eq AS eq, 
			sold AS sold, 
			dhl AS dhl, 
			ar AS ar, 
			sbl AS sbl, 
			pmf AS pmf, 
			pmm AS pmm, 
			ticket AS ticket, 
			mtc AS mtc, 
			ws AS ws, 
			mast AS mast, 
			oris AS oris, 
			elite AS elite, 
			yes AS yes, 
			time AS time, 
			comment AS comment,  
			provinces.name AS province, 
			sups.name AS sup, 
			users.fullname AS user, 
			areas.name AS area,
			pos.shop AS pos
		`).
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
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
		"message":    "All posform by province",
		"data":       dataList,
		"pagination": pagination,
	})
}

func GetPosformBySupID(c *fiber.Ctx) error {
	db := database.DB
	areaId := c.Params("area_id")

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

	var dataList []models.PosForm

	var length int64
	var data []models.PosForm
	db.Model(data).Where("area_id = ?", areaId).Count(&length)

	db.
		Joins("JOIN provinces ON pos_forms.province_id=provinces.id").
		Joins("JOIN sups ON pos_forms.sup_id=sups.id").
		Joins("JOIN users ON pos_forms.user_id=users.id").
		Joins("JOIN areas ON pos_forms.area_id=areas.id").
		Joins("JOIN pos ON pos_forms.pos_id=pos.id").
		Where("pos_forms.area_id = ?", areaId).
		Where("users.fullname ILIKE ? OR provinces.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
			pos_forms.id AS id,
			id_unique AS id_unique,
			eq AS eq, 
			sold AS sold, 
			dhl AS dhl, 
			ar AS ar, 
			sbl AS sbl, 
			pmf AS pmf, 
			pmm AS pmm, 
			ticket AS ticket, 
			mtc AS mtc, 
			ws AS ws, 
			mast AS mast, 
			oris AS oris, 
			elite AS elite, 
			yes AS yes, 
			time AS time, 
			comment AS comment,  
			provinces.name AS province, 
			sups.name AS sup, 
			users.fullname AS user, 
			areas.name AS area,
			pos.shop AS pos
		`).
		Offset(offset).
		Limit(limit).
		Order("pos_forms.updated_at DESC").
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
		"message":    "All posform",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllPosforms(c *fiber.Ctx) error {
	db := database.DB
	var data []models.PosForm
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All PosForms",
		"data":    data,
	})
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

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "posForm created success",
			"data":    p,
		},
	)
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
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`

		Price string `json:"price"`
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

	posform.Latitude = updateData.Latitude
	posform.Longitude = updateData.Longitude
	posform.Price = updateData.Price

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
