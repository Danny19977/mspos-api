package poss

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedPos(c *fiber.Ctx) error {
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

	var u []models.Pos
	var length int64
	db := database.DB
	db.Find(&u).Count(&length)

	sql1 := `
		SELECT "pos"."id" AS id, 
		status AS status, 
		"pos"."name" AS name, 
		"pos"."shop" AS shop,
		"pos"."manager" AS manager,  
		"pos"."telephone" AS telephone,
		"provinces"."name" AS province,   
		"areas"."name" AS area,
		"pos"."commune" AS commune,  
		"pos"."quartier" AS quartier,  
		"pos"."avenue" AS avenue,  
		"pos"."reference" AS reference,
		
		"pos"."eparasol" AS eparasol,
		"pos"."etable" AS etable,
		"pos"."ekiosk" AS ekiosk,
		"pos"."input_group_selector" AS input_group_selector,
		"pos"."cparasol" AS cparasol,
		"pos"."ctable" AS ctable,
		"pos"."ckiosk" AS ckiosk 
		FROM pos 
			INNER JOIN provinces ON pos.province_id=provinces.id  
			INNER JOIN areas ON pos.area_id=areas.id  
			WHERE "pos"."deleted_at" IS NULL
		ORDER BY "pos"."updated_at" DESC;
	`
	var dataList []models.PosPaginate
	database.DB.Raw(sql1).Scan(&dataList)

	if offset >= len(dataList) {
		dataList = []models.PosPaginate{} // Empty slice
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
