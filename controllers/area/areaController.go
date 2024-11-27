package area

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedAreas(c *fiber.Ctx) error {
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

	var dataList []models.Area

	var length int64
	// var data []models.Sup
	db.Model(dataList).Count(&length)

	db.
		Joins("JOIN provinces ON areas.province_id=provinces.id").
		Joins("JOIN sups ON areas.sup_id=sups.id").
		Where("areas.name ILIKE ?", "%"+search+"%").
		Select(`
		areas.id AS id,  
		areas.name AS name,
		provinces.name AS province, 
		areas.commune AS commune,
		sups.name AS sup
	`).
		Offset(offset).
		Limit(limit).
		Order("areas.updated_at DESC").
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
		"message":    "All areas",
		"data":       dataList,
		"pagination": pagination,
	})
}

// query data province
func GetAreaByProvinceID(c *fiber.Ctx) error {
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

	var dataList []models.Area	

	var length int64
	// var data []models.Area
	db.Model(dataList).Where("province_id = ?", provinceId).Count(&length) 

	db.
		Joins("JOIN provinces ON areas.province_id=provinces.id").
		Joins("JOIN sups ON areas.sup_id=sups.id").
		Where("areas.province_id = ?", provinceId).
		Where("areas.name ILIKE ?", "%"+search+"%").
		Select(`
		areas.id AS id,  
		areas.name AS name,
		provinces.name AS province, 
		areas.commune AS commune,
		sups.name AS sup
	`).
		Offset(offset).
		Limit(limit).
		Order("areas.updated_at DESC").
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
		"message":    "All areas",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All Provinces Dropdown
func GetAreaDropdown(c *fiber.Ctx) error {
	db := database.DB
	sql1 := `
		SELECT "areas"."id" AS id, 
			"areas"."name" AS name, 
			"areas"."province_id" AS province_id,
			"areas"."commune" AS commune
		FROM pos_forms
		INNER JOIN areas ON pos_forms.area_id=areas.id
		WHERE "pos_forms"."deleted_at" IS NULL
		GROUP BY "areas"."id", "areas"."name", "areas"."province_id";
	`
	var data []models.AreaDropDown
	db.Raw(sql1).Scan(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All area Dropdown",
		"data":    data,
	})
}

// Get All data
func GetAllAreas(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Area
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Areas",
		"data":    data,
	})
}

// query data
func GetAreaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var areas []models.Area
	db.Where("province_id = ?", id).Find(&areas)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "areas by id found",
		"data":    areas,
	})
}

// query data
func GetSupAreaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var areas []models.Area
	db.Where("sup_id = ?", id).Find(&areas)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "poss by id found",
		"data":    areas,
	})
}

// Get one data
func GetArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var area models.Area
	db.Find(&area, id)
	if area.Name == "" {
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

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "area created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name       string `json:"name"`
		ProvinceID uint   `json:"province_id"`
		SupID      uint   `json:"sup_id"`
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

	area := new(models.Area)

	db.First(&area, id)
	area.Name = updateData.Name
	area.ProvinceID = updateData.ProvinceID
	area.SupID = updateData.SupID
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
	if area.Name == "" {
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
