package PosFormItem

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Paginate
func GetPaginatedStock(c *fiber.Ctx) error {
	db := database.DB
	posformID := c.Params("posform_id")

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit
	
	// Deferent filter 
	search := c.Query("search", "")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var dataList []models.PosFormItems
	var length int64
	db.Model(dataList).Where("posform_id = ?", posformID).Count(&length)
	db.Where("posform_id = ?", posformID).
		Joins("JOIN posforms ON posform_item.posform_id=posforms.id").
		Joins("JOIN brands ON posform_item.posform_id=brands.id"). 
		Where("posform_item.created_at BETWEEN ? AND ?", start_date, end_date).
		Where("brands.name ILIKE ? OR posforms.reference ILIKE ? OR brands.name ILIKE ?", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("stocks.created_at DESC").
		
		Preload("Product").
		Preload("Fournisseur").
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
		"message":    "All stocks",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get data
func GetStockMargeBeneficiaire(c *fiber.Ctx) error {
	db := database.DB
	productId := c.Params("product_id")

	var data models.PosFormItems 

	db.Model(data).Where("product_id = ?", productId).Preload("Product").Last(&data)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Total qty stocks",
		"data":    data,
	})
}


// Get Total data
func GetTotalStock(c *fiber.Ctx) error {
	db := database.DB
	productId := c.Params("product_id")

	var data []models.PosFormItems
	var totalQty int64

	db.Model(data).Where("product_id = ?", productId).Select("SUM(quantity)").Scan(&totalQty)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Total qty stocks",
		"data":    totalQty,
	})
}

// Get All data
func GetAllStocks(c *fiber.Ctx) error {
	db := database.DB
	productId := c.Params("product_id")
	var data []models.PosFormItems
	db.Where("product_id = ?", productId).Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All stocks",
		"data":    data,
	})
}


// Get one data
func GetStock(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var stock models.PosFormItems
	db.Find(&stock, id)
	if stock.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No stock name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock found",
			"data":    stock,
		},
	)
}

// Create data
func CreateStock(c *fiber.Ctx) error {
	p := &models.PosFormItems{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateStock(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		PosID          uint      `json:"pos_id"`
		ProductID      uint      `json:"product_id"`
		Description    string    `json:"description"`
		FournisseurID  uint      `json:"fournisseur_id"`
		Quantity       uint64    `json:"quantity"`
		PrixAchat      float64   `json:"prix_achat"`
		DateExpiration time.Time `json:"date_expiration"`
		Signature      string    `json:"signature"`
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

	stock := new(models.PosFormItems)

	db.First(&stock, id)
	stock.PosID = updateData.PosID
	stock.ProductID = updateData.ProductID
	stock.Description = updateData.Description
	stock.FournisseurID = updateData.FournisseurID
	stock.Quantity = updateData.Quantity
	stock.PrixAchat = updateData.PrixAchat
	stock.DateExpiration = updateData.DateExpiration
	stock.Signature = updateData.Signature

	db.Save(&stock)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock updated success",
			"data":    stock,
		},
	)

}

// Delete data
func DeleteStock(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var stock models.PosFormItems
	db.First(&stock, id)
	if stock.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No stock name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&stock)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock deleted success",
			"data":    nil,
		},
	)
}