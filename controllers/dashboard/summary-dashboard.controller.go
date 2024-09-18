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

func TrackingVisitDRS(c *fiber.Ctx) error {
	days := c.Params("days")
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")


	sql1 := `
		
	SELECT "provinces"."name" AS province,
		ROUND(SUM(eq1) / COUNT("pos_forms"."id") * 100) AS nd,
		ROUND(SUM(eq) / (SUM(eq) + SUM(dhl) + SUM(ar) +
		SUM(sbl) +
		SUM(pmf) +
		SUM(pmm) +
		SUM(ticket) +
		SUM(mtc) +
		SUM(ws) +
		SUM(mast) +
		SUM(oris) +
		SUM(elite) +
		SUM(yes) +
		SUM(time) ) * 100) AS sos,
	ROUND(100 - ROUND(SUM(eq1) / COUNT("pos_forms"."id") * 100)) AS oos,
	(SELECT COUNT(*) FROM users 
		INNER JOIN provinces ON users.province_id=provinces.id
		WHERE role='DR' AND status=true AND province_id="provinces"."id") AS dr,
		
	COUNT("pos_forms"."id") AS visit,

	ROUND(40 * (SELECT COUNT(*) 
			FROM users 
			INNER JOIN provinces ON users.province_id=provinces.id
				WHERE role='DR' AND status=true AND province_id="provinces"."id") * 
	CASE 
		WHEN ? = 0 THEN 1
			ELSE ?
		END ) AS obj,
	
		ROUND(COUNT("pos_forms"."id") / (40 * (SELECT COUNT(*) 
			FROM users 
			INNER JOIN provinces ON users.province_id=provinces.id
				WHERE role='DR' AND status=true AND province_id="provinces"."id") * 
	CASE 
			WHEN ? = 0 THEN 1
			ELSE ?
		END )) AS perf
	
			FROM pos_forms 
					INNER JOIN provinces ON pos_forms.province_id=provinces.id
				INNER JOIN users ON pos_forms.user_id=users.id
					WHERE "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP  
					AND ? ::TIMESTAMP

	GROUP BY "provinces"."name";
	`
	var chartData []models.TrackingVisitDRSChart
	database.DB.Raw(sql1, days, days, days, days, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}


func SummaryChartBar(c *fiber.Ctx) error { 
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		SELECT "provinces"."name" AS province,
			ROUND(SUM(eq1) / COUNT("pos_forms"."id") * 100) AS nd,
			ROUND(SUM(eq) / (SUM(eq) + SUM(dhl) + SUM(ar) +
			SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) +
			SUM(mtc) + SUM(ws) + SUM(mast) +
			SUM(oris) + SUM(elite) + SUM(yes) + SUM(time) 
		) * 100) AS sos,
		ROUND(100 - ROUND(SUM(eq1) / COUNT("pos_forms"."id") * 100)) AS oos
		FROM pos_forms 
				INNER JOIN provinces ON pos_forms.province_id=provinces.id
			INNER JOIN users ON pos_forms.user_id=users.id
				WHERE "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP  
				AND ? ::TIMESTAMP

		GROUP BY "provinces"."name";
	`
	var chartData []models.SumChartBar
	database.DB.Raw(sql1, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}
