package posform

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedPosForm(c *fiber.Ctx) error {
	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page") // CurrentPage

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 15
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	offset := (page - 1) * pageSize

	var u []models.PosForm
	var length int64
	db := database.DB
	db.Find(&u).Count(&length)

	sql1 := `
		SELECT "pos_forms"."id" AS id, 
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
		"provinces"."name" AS province, 
		"sups"."name" AS sup, 
		"users"."fullname" AS user, 
		"areas"."name" AS area,
		"pos"."shop" AS pos
		FROM pos_forms 
			INNER JOIN provinces ON pos_forms.province_id=provinces.id 
			INNER JOIN sups ON pos_forms.sup_id=sups.id 
			INNER JOIN users ON pos_forms.user_id=users.id 
			INNER JOIN areas ON pos_forms.area_id=areas.id 
			INNER JOIN pos ON pos_forms.pos_id=pos.id 
 WHERE "pos_forms"."deleted_at" IS NULL 
		ORDER BY "pos_forms"."updated_at" DESC;
	`
	var dataList []models.PosFormPaginate
	database.DB.Raw(sql1).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.PosFormPaginate{} // Empty slice
	} else {
		end := offset + pageSize
		if end > len(dataList) {
			end = len(dataList)
		}
		dataList = dataList[offset:end]
	}
	// Calculate total number of pages
	totalPages := len(dataList) / pageSize
	if remainder := len(dataList) % pageSize; remainder > 0 {
		totalPages++
	}

	// Create pagination metadata (adjust fields as needed)
	pagination := map[string]interface{}{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   pageSize,
		"length":      length,
	}

	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All PosForms",
		"data":       dataList,
		"pagination": pagination,
	})
}

// query data dr
func GetPosformByID(c *fiber.Ctx) error {
	userId := c.Params("id")

	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page") // CurrentPage

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 15
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	offset := (page - 1) * pageSize

	var u []models.PosForm
	var length int64
	db := database.DB
	db.Where("user_id = ?", userId).Find(&u).Count(&length)

	sql1 := `
	SELECT "pos_forms"."id" AS id, 
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
		"provinces"."name" AS province, 
		"sups"."name" AS sup, 
		"users"."fullname" AS user, 
		"areas"."name" AS area,
		"pos"."shop" AS pos
		FROM pos_forms 
			INNER JOIN provinces ON pos_forms.province_id=provinces.id 
			INNER JOIN sups ON pos_forms.sup_id=sups.id 
			INNER JOIN users ON pos_forms.user_id=users.id 
			INNER JOIN areas ON pos_forms.area_id=areas.id 
			INNER JOIN pos ON pos_forms.pos_id=pos.id 
 		WHERE "pos_forms"."deleted_at" IS NULL AND "pos_forms"."user_id"=?
		ORDER BY "pos_forms"."updated_at" DESC; 
	`
	var dataList []models.PosFormPaginate
	database.DB.Raw(sql1, userId).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.PosFormPaginate{} // Empty slice
	} else {
		end := offset + pageSize
		if end > len(dataList) {
			end = len(dataList)
		}
		dataList = dataList[offset:end]
	}
	// Calculate total number of pages
	totalPages := len(dataList) / pageSize
	if remainder := len(dataList) % pageSize; remainder > 0 {
		totalPages++
	}

	// Create pagination metadata (adjust fields as needed)
	pagination := map[string]interface{}{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   pageSize,
		"length":      length,
	}

	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All posform by dr",
		"data":       dataList,
		"pagination": pagination,
	})
}

// query data province
func GetPosformByProvinceID(c *fiber.Ctx) error {
	provinceId := c.Params("id")

	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page") // CurrentPage

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 15
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	offset := (page - 1) * pageSize

	var u []models.PosForm
	var length int64
	db := database.DB
	db.Where("province_id = ?", provinceId).Find(&u).Count(&length)

	sql1 := `
	SELECT "pos_forms"."id" AS id, 
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
		"provinces"."name" AS province, 
		"sups"."name" AS sup, 
		"users"."fullname" AS user, 
		"areas"."name" AS area,
		"pos"."shop" AS pos
		FROM pos_forms 
			INNER JOIN provinces ON pos_forms.province_id=provinces.id 
			INNER JOIN sups ON pos_forms.sup_id=sups.id 
			INNER JOIN users ON pos_forms.user_id=users.id 
			INNER JOIN areas ON pos_forms.area_id=areas.id 
			INNER JOIN pos ON pos_forms.pos_id=pos.id 
 		WHERE "pos_forms"."deleted_at" IS NULL AND "pos_forms"."province_id"=?
		ORDER BY "pos_forms"."updated_at" DESC;
	`
	var dataList []models.PosFormPaginate
	database.DB.Raw(sql1, provinceId).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.PosFormPaginate{} // Empty slice
	} else {
		end := offset + pageSize
		if end > len(dataList) {
			end = len(dataList)
		}
		dataList = dataList[offset:end]
	}
	// Calculate total number of pages
	totalPages := len(dataList) / pageSize
	if remainder := len(dataList) % pageSize; remainder > 0 {
		totalPages++
	}

	// Create pagination metadata (adjust fields as needed)
	pagination := map[string]interface{}{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   pageSize,
		"length":      length,
	}

	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All posform by province",
		"data":       dataList,
		"pagination": pagination,
	})
}

// query data sup
func GetPosformBySupID(c *fiber.Ctx) error {
	areaId := c.Params("id")

	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page") // CurrentPage

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 15
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	offset := (page - 1) * pageSize

	var u []models.PosForm
	var length int64
	db := database.DB
	db.Where("area_id = ?", areaId).Find(&u).Count(&length)

	sql1 := `
	SELECT "pos_forms"."id" AS id, 
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
		"provinces"."name" AS province, 
		"sups"."name" AS sup, 
		"users"."fullname" AS user, 
		"areas"."name" AS area,
		"pos"."shop" AS pos
		FROM pos_forms 
			INNER JOIN provinces ON pos_forms.province_id=provinces.id 
			INNER JOIN sups ON pos_forms.sup_id=sups.id 
			INNER JOIN users ON pos_forms.user_id=users.id 
			INNER JOIN areas ON pos_forms.area_id=areas.id 
			INNER JOIN pos ON pos_forms.pos_id=pos.id 
 		WHERE "pos_forms"."deleted_at" IS NULL AND "pos_forms"."area_id"=?
		ORDER BY "pos_forms"."updated_at" DESC;
	`
	var dataList []models.PosFormPaginate
	database.DB.Raw(sql1, areaId).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.PosFormPaginate{} // Empty slice
	} else {
		end := offset + pageSize
		if end > len(dataList) {
			end = len(dataList)
		}
		dataList = dataList[offset:end]
	}
	// Calculate total number of pages
	totalPages := len(dataList) / pageSize
	if remainder := len(dataList) % pageSize; remainder > 0 {
		totalPages++
	}

	// Create pagination metadata (adjust fields as needed)
	pagination := map[string]interface{}{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   pageSize,
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
		Eq        int64   `json:"eq"`
		Eq1       int64   `json:"eq1"`
		Sold      int64   `json:"sold"`
		Dhl       int64   `json:"dhl"`
		Dhl1      int64   `json:"dhl1"`
		Ar        int64   `json:"ar"`
		Ar1       int64   `json:"ar1"`
		Sbl       int64   `json:"sbl"`
		Sbl1      int64   `json:"sbl1"`
		Pmf       int64   `json:"pmf"`
		Pmf1      int64   `json:"pmf1"`
		Pmm       int64   `json:"pmm"`
		Pmm1      int64   `json:"pmm1"`
		Ticket    int64   `json:"ticket"`
		Ticket1   int64   `json:"ticket1"`
		Mtc       int64   `json:"mtc"`
		Mtc1      int64   `json:"mtc1"`
		Ws        int64   `json:"ws"`
		Ws1       int64   `json:"ws1"`
		Mast      int64   `json:"mast"`
		Mast1     int64   `json:"mast1"`
		Oris      int64   `json:"oris"`
		Oris1     int64   `json:"oris1"`
		Elite     int64   `json:"elite"`
		Elite1    int64   `json:"elite1"`
		Yes       int64   `json:"yes"`
		Yes1      int64   `json:"yes1"`
		Time      int64   `json:"time"`
		Time1     int64   `json:"time1"`
		Comment   string  `json:"comment"`
		Signature string  `json:"signature"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`

		Price int64 `json:"price"`
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
