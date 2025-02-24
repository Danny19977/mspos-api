package routeplan

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
	// "github.com/lib/pq"
)

// Paginate
func GetPaginatedRouthplan(c *fiber.Ctx) error {
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

	var dataList []models.RoutePlan

	var length int64
	db.Model(dataList).Count(&length)
	db.Joins("User").
		Joins("Province").
		Joins("Area").
		Joins("SubArea").
		Where("route_plans.name ILIKE ? OR users.name ILIKE ? OR provinces.name ILIKE ? OR areas.name ILIKE ? OR sub_areas.name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("route_plans.updated_at DESC").
		Preload("User").
		Preload("Province").
		Preload("Area").
		Preload("SubArea").
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
		"message":    "All products",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllRouthplan(c *fiber.Ctx) error {
	db := database.DB

	var data []models.SubArea
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Routhplan",
		"data":    data,
	})
}

// Get All data by id
func GetAllRouthplanBySearch(c *fiber.Ctx) error {
	db := database.DB

	search := c.Query("search", "")

	var data []models.RoutePlan
	db.Where("name ILIKE ?", "%"+search+"%").
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Routhplan",
		"data":    data,
	})
}

// Get one data
func GetRouthplan(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var Routeplan models.RoutePlan
	db.Find(&Routeplan, id)
	if Routeplan.ID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No Routeplan name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "subarea found",
			"data":    Routeplan,
		},
	)
}

// Create data
func CreateRouthplan(c *fiber.Ctx) error {
	p := &models.RoutePlan{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "routeplan created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateRouthplan(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name       string         `json:"name"`
		
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

	subarea := new(models.SubArea)

	db.First(&subarea, id)
	subarea.Name = updateData.Name

	db.Save(&subarea)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "subarea updated success",
			"data":    subarea,
		},
	)

}

// Delete data
func DeleteRouthplan(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var routeplan models.RoutePlan
	db.First(&routeplan, id)
	if routeplan.ID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No routeplan name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&routeplan)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "subarea deleted success",
			"data":    nil,
		},
	)
}
