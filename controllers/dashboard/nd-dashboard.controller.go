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
		(SUM(eq1) / COUNT("pos_forms"."id")) * 100 AS eq,
		(SUM(dhl1) / COUNT("pos_forms"."id")) * 100 AS dhl, 
		(SUM(ar1) / COUNT("pos_forms"."id")) * 100 AS ar, 
		(SUM(sbl1)/ COUNT("pos_forms"."id")) * 100 AS sbl, 
		(SUM(pmf1) / COUNT("pos_forms"."id")) * 100 AS pmf,
		(SUM(pmm1) / COUNT("pos_forms"."id")) * 100 AS pmm, 
		(SUM(ticket1) / COUNT("pos_forms"."id")) * 100 AS ticket, 
		(SUM(mtc1) / COUNT("pos_forms"."id")) * 100 AS mtc, 
		(SUM(ws1) / COUNT("pos_forms"."id")) * 100 AS ws, 
		(SUM(mast1) / COUNT("pos_forms"."id")) * 100 AS mast,
		(SUM(oris1) / COUNT("pos_forms"."id")) * 100 AS oris, 
		(SUM(elite1) / COUNT("pos_forms"."id")) * 100 AS elite,
		(SUM(yes1) / COUNT("pos_forms"."id")) * 100 AS yes, 
		(SUM(time1) / COUNT("pos_forms"."id")) * 100 AS time
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
	SELECT EXTRACT(MONTH FROM "pos_forms"."created_at") AS month,
		CAST(SUM(eq1) / COUNT(*) as decimal(40,0)) * 100 AS eq
	FROM pos_forms
	INNER JOIN provinces ON pos_forms.province_id=provinces.id
		WHERE "provinces"."name"=? AND EXTRACT(YEAR FROM "pos_forms"."created_at") = EXTRACT(YEAR FROM CURRENT_DATE)
		AND EXTRACT(MONTH FROM "pos_forms"."created_at") BETWEEN 1 AND 12 
		GROUP BY EXTRACT(MONTH FROM "pos_forms"."created_at")
		ORDER BY EXTRACT(MONTH FROM "pos_forms"."created_at");
	`
	var chartData  []models.NdByYear
	database.DB.Raw(sql1, province).Scan(&chartData)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}
