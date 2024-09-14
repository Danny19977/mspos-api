package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

func NdTableView(c *fiber.Ctx) error {
	province := c.Params("province")
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
	SELECT areas.name AS area, 
		COUNT(*) AS total,
		CAST(SUM(eq1) / COUNT(*) as decimal(40,0)) * 100 AS eq,
		CAST(SUM(dhl1) / COUNT(*) as decimal(40,0)) * 100 AS dhl, 
		CAST(SUM(ar1) / COUNT(*) as decimal(40,0)) * 100 AS ar, 
		CAST(SUM(sbl1)/ COUNT(*) as decimal(40,0)) * 100 AS sbl, 
		CAST(SUM(pmf1) / COUNT(*) as decimal(40,0)) * 100 AS pmf,
		CAST(SUM(pmm1) / COUNT(*) as decimal(40,0)) * 100 AS pmm, 
		CAST( SUM(ticket1) / COUNT(*) as decimal(40,0)) * 100 AS ticket, 
		CAST( SUM(mtc1) / COUNT(*) as decimal(40,0)) * 100 AS mtc, 
		CAST( SUM(ws1) / COUNT(*) as decimal(40,0)) * 100 AS ws, 
		CAST( SUM(mast1) / COUNT(*) as decimal(40,0)) * 100 AS mast,
		CAST(SUM(oris1) / COUNT(*) as decimal(40,0)) * 100 AS oris, 
		CAST( SUM(elite1) / COUNT(*) as decimal(40,0)) * 100 AS elite, 
		CAST(SUM(yes1) / COUNT(*) as decimal(40,0)) * 100 AS yes, 
		CAST( SUM(time1) / COUNT(*) as decimal(40,0)) * 100 AS time
	FROM pos_forms
	INNER JOIN areas ON pos_forms.area_id=areas.id
	INNER JOIN provinces ON pos_forms.province_id=provinces.id
	WHERE "provinces"."name"=? AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP 
		AND ? ::TIMESTAMP 
	GROUP BY areas.name;
	`
	var chartData []models.NDChartData
	database.DB.Raw(sql1, province, start_date, end_date).Scan(&chartData)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}

 
func NdByYear(c *fiber.Ctx) error {
	province := c.Params("province") 

	sql1 := `  
	SELECT EXTRACT(MONTH FROM "pos_forms"."created_at") AS mois,
		CAST(SUM(eq1) / COUNT(*) as decimal(40,0)) * 100 AS eq
		FROM pos_forms
		INNER JOIN provinces ON pos_forms.province_id=provinces.id
				WHERE "provinces"."name"=? AND EXTRACT(YEAR FROM "pos_forms"."created_at") = EXTRACT(YEAR FROM CURRENT_DATE)
		AND EXTRACT(MONTH FROM "pos_forms"."created_at") BETWEEN 1 AND 12 
	GROUP BY mois
	ORDER BY mois;
	`
	var chartData models.NdByYear
	database.DB.Raw(sql1, province).Scan(&chartData)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}
