package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

func DrCount(c *fiber.Ctx) error {
	sql1 := `
	 SELECT COUNT(*) FROM users WHERE role='DR' AND status=true;
	`
	var chartData models.SummaryCount
	database.DB.Raw(sql1).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func POSCount(c *fiber.Ctx) error {
	sql1 := `
	 SELECT COUNT(*) FROM pos WHERE status=true;
	`

	var chartData models.SummaryCount
	database.DB.Raw(sql1).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func ProvinceCount(c *fiber.Ctx) error {
	sql1 := `
	 SELECT COUNT(*) FROM provinces;
	`

	var chartData models.SummaryCount
	database.DB.Raw(sql1).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

func AreaCount(c *fiber.Ctx) error {
	sql1 := `
	 SELECT COUNT(*) FROM areas;
	`

	var chartData models.SummaryCount
	database.DB.Raw(sql1).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}


func SOSPie(c *fiber.Ctx) error {
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")
	
	sql1 := `
		SELECT "provinces"."name" AS province,
			ROUND(SUM(eq) / (SUM(eq) + SUM(dhl) + SUM(ar) +
			SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
			SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
			SUM(time) ) * 100) AS eq
		FROM pos_forms 
			INNER JOIN provinces ON pos_forms.province_id=provinces.id
				WHERE "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP  
				AND ? ::TIMESTAMP 
		GROUP BY "provinces"."name"; 
	`

	var chartData []models.SosPieChart
	database.DB.Raw(sql1, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}