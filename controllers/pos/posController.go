package poss

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedPos(c *fiber.Ctx) error {
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

	var dataList []models.Pos

	var length int64
	// var data []models.Pos
	db.Model(dataList).Count(&length)

	db.
		Joins("JOIN provinces ON pos.province_id=provinces.id").
		Joins("JOIN users ON pos.user_id=users.id").
		Joins("JOIN areas ON pos.area_id=areas.id").
		Where("users.fullname ILIKE ? OR pos.name ILIKE ? OR pos.shop ILIKE ? OR pos.manager ILIKE ? OR pos.telephone ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
			pos.id AS id, 
			pos.status AS status, 
			provinces.name AS province,
			areas.name AS area,
			users.fullname AS dr,
			pos.name AS name, 
			pos.shop AS shop,
			pos.manager AS manager,  
			pos.telephone AS telephone, 
			pos.commune AS commune,  
			pos.quartier AS quartier,  
			pos.avenue AS avenue,  
			pos.reference AS reference,
			pos.eparasol AS eparasol,
			pos.etable AS etable,
			pos.ekiosk AS ekiosk,
			pos.input_group_selector AS input_group_selector,
			pos.cparasol AS cparasol,
			pos.ctable AS ctable,
			pos.ckiosk AS ckiosk
		`).
		Offset(offset).
		Limit(limit).
		Order("pos.updated_at DESC").
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
		"message":    "All Pos",
		"data":       dataList,
		"pagination": pagination,
	})
}

// query data DR
func GetPosPaginateByID(c *fiber.Ctx) error {
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

	var dataList []models.Pos

	var length int64
	// var data []models.Pos
	db.Model(dataList).Where("user_id = ?", userId).Count(&length)

	db.
		Joins("JOIN provinces ON pos.province_id=provinces.id").
		Joins("JOIN users ON pos.user_id=users.id").
		Joins("JOIN areas ON pos.area_id=areas.id").
		Where("pos.user_id = ?", userId).
		Where("users.fullname ILIKE ? OR pos.name ILIKE ? OR pos.shop ILIKE ? OR pos.manager ILIKE ? OR pos.telephone ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
		pos.id AS id,
		pos.status AS status, 
		provinces.name AS province,
		areas.name AS area,
		users.fullname AS dr,
		pos.name AS name, 
		pos.shop AS shop,
		pos.manager AS manager,  
		pos.telephone AS telephone, 
		pos.commune AS commune,  
		pos.quartier AS quartier,  
		pos.avenue AS avenue,  
		pos.reference AS reference,
		pos.eparasol AS eparasol,
		pos.etable AS etable,
		pos.ekiosk AS ekiosk,
		pos.input_group_selector AS input_group_selector,
		pos.cparasol AS cparasol,
		pos.ctable AS ctable,
		pos.ckiosk AS ckiosk 
	`).
		Offset(offset).
		Limit(limit).
		Order("pos.updated_at DESC").
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
		"message":    "All pos by DR",
		"data":       dataList,
		"pagination": pagination,
	})
}

// query data province
func GetPosByProvinceID(c *fiber.Ctx) error {
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

	var dataList []models.Pos

	var length int64
	
	db.Model(dataList).Where("province_id = ?", provinceId).Count(&length)

	db.
		Joins("JOIN provinces ON pos.province_id=provinces.id").
		Joins("JOIN users ON pos.user_id=users.id").
		Joins("JOIN areas ON pos.area_id=areas.id").
		Where("pos.province_id = ?", provinceId).
		Where("users.fullname ILIKE ? OR pos.name ILIKE ? OR pos.shop ILIKE ? OR pos.manager ILIKE ? OR pos.telephone ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
			pos.id AS id, 
			pos.status AS status, 
			provinces.name AS province,
			areas.name AS area,
			users.fullname AS dr,
			pos.name AS name, 
			pos.shop AS shop,
			pos.manager AS manager,  
			pos.telephone AS telephone, 
			pos.commune AS commune,  
			pos.quartier AS quartier,  
			pos.avenue AS avenue,  
			pos.reference AS reference,
			pos.eparasol AS eparasol,
			pos.etable AS etable,
			pos.ekiosk AS ekiosk,
			pos.input_group_selector AS input_group_selector,
			pos.cparasol AS cparasol,
			pos.ctable AS ctable,
			pos.ckiosk AS ckiosk 
		`).
		Offset(offset).
		Limit(limit).
		Order("pos.updated_at DESC").
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
		"message":    "All pos by province",
		"data":       dataList,
		"pagination": pagination,
	})
}

// query data sup by area
func GetPosBySupID(c *fiber.Ctx) error {
	db := database.DB
	areaId := c.Params("sup_id")

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

	var dataList []models.Pos

	var length int64
	// var data []models.Pos
	db.Model(dataList).Where("area_id = ?", areaId).Count(&length)
	db.
		Joins("JOIN provinces ON pos.province_id=provinces.id").
		Joins("JOIN users ON pos.user_id=users.id").
		Joins("JOIN areas ON pos.area_id=areas.id").
		Joins("JOIN pos ON pos.pos_id=pos.id").
		Where("pos.area_id = ?", areaId).
		Where("users.fullname ILIKE ? OR pos.name ILIKE ? OR pos.shop ILIKE ? OR pos.manager ILIKE ? OR pos.telephone ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Select(`
			pos.id AS id, 
			pos.status AS status, 
			provinces.name AS province,
			areas.name AS area,
			users.fullname AS dr,
			pos.name AS name, 
			pos.shop AS shop,
			pos.manager AS manager,  
			pos.telephone AS telephone, 
			pos.commune AS commune,  
			pos.quartier AS quartier,  
			pos.avenue AS avenue,  
			pos.reference AS reference,
			pos.eparasol AS eparasol,
			pos.etable AS etable,
			pos.ekiosk AS ekiosk,
			pos.input_group_selector AS input_group_selector,
			pos.cparasol AS cparasol,
			pos.ctable AS ctable,
			pos.ckiosk AS ckiosk 
		`).
		Offset(offset).
		Limit(limit).
		Order("pos.updated_at DESC").
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
		"message":    "All pos by area",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get all data by search
func GetAllPosSearch(c *fiber.Ctx) error {
	db := database.DB
	search := c.Query("search", "")

	var data []models.Pos
	if search != "" {
		db.Where("name ILIKE ? OR shop ILIKE ?", "%"+search+"%", "%"+search+"%").Find(&data)
	} else {
		db.Find(&data)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// Get All data
func GetAllPoss(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Pos
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Pos",
		"data":    data,
	})
}

// query data
func GetPosByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var poss []models.Pos
	db.Where("province_id = ?", id).Find(&poss)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "poss by id found",
		"data":    poss,
	})
}

// query data Area
func GetPosAreaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var poss []models.Pos
	db.Where("area_id = ?", id).Find(&poss)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "poss by id found",
		"data":    poss,
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

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "pos created success",
			"data":    p,
		},
	)
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
		Telephone          string `json:"telephone"`
		Eparasol           bool   `json:"eparasol"`
		Etable             bool   `json:"etable"`
		Ekiosk             bool   `json:"ekiosk"`
		InputGroupSelector string `json:"inputgroupselector"`
		Cparasol           bool   `json:"cparasol"`
		Ctable             bool   `json:"ctable"`
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
