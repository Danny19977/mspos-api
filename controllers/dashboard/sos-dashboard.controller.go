package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

func SOSPieByArea(c *fiber.Ctx) error {
	province := c.Params("province")
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
		SELECT "areas"."name" AS area,
			ROUND(SUM(eq) / (SUM(eq) + SUM(dhl) + SUM(ar) +
			SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
			SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
			SUM(time) ) * 100) AS eq
		FROM pos_forms 
			INNER JOIN provinces ON pos_forms.province_id=provinces.id
			INNER JOIN areas ON pos_forms.area_id=areas.id
		WHERE "provinces"."name"=? AND 
				"pos_forms"."created_at" BETWEEN ? ::TIMESTAMP AND ? ::TIMESTAMP 
		GROUP BY "areas"."name";
	`

	var chartData []models.SosPieChartArea
	database.DB.Raw(sql1, province, start_date, end_date).Scan(&chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}



func SOSByYear(c *fiber.Ctx) error {
	province := c.Params("province")
	sql1 := `  
	SELECT EXTRACT(MONTH FROM "pos_forms"."created_at") AS month,
		ROUND(SUM(eq) / (SUM(eq) + SUM(dhl) + SUM(ar) +
			SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
			SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
			SUM(time) ) * 100) AS eq
	FROM pos_forms
	INNER JOIN provinces ON pos_forms.province_id=provinces.id
	WHERE "provinces"."name"=? AND 
    EXTRACT(YEAR FROM "pos_forms"."created_at") = EXTRACT(YEAR FROM CURRENT_DATE)
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

func SOSTableView(c *fiber.Ctx) error {
	province := c.Params("province")
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := `
	SELECT areas.name AS area, 
		ROUND(SUM(eq) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS eq,
		ROUND(SUM(dhl) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS dhl, 
		ROUND(SUM(ar) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS ar, 
		ROUND(SUM(sbl) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS sbl, 
		ROUND(SUM(pmf) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS pmf, 
		ROUND(SUM(pmm) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS pmm, 
		ROUND(SUM(ticket) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS ticket, 
		ROUND(SUM(mtc) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS mtc, 
		ROUND(SUM(ws) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS ws, 
		ROUND(SUM(mast) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS mast, 
		ROUND(SUM(oris) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS oris, 
		ROUND(SUM(elite) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS elite,
		ROUND(SUM(yes) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS yes,
		ROUND(SUM(time) / (SUM(eq) + SUM(dhl) + SUM(ar) +
				SUM(sbl) + SUM(pmf) + SUM(pmm) + SUM(ticket) + SUM(mtc) +
				SUM(ws) + SUM(mast) + SUM(oris) + SUM(elite) + SUM(yes) +
				SUM(time) ) * 100) AS time
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