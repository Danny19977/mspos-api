package dashboard

import (
	"fmt"

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
		CAST((SELECT COUNT(equateur) FROM pos_forms WHERE equateur::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as eq, 
		CAST((SELECT COUNT(dhl) FROM pos_forms WHERE dhl::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as dhl,
		CAST((SELECT COUNT(ar) FROM pos_forms WHERE ar::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float  as decimal(40,1))* 100 as ar,
		CAST((SELECT COUNT(sbl) FROM pos_forms WHERE sbl::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(20,1)) * 100 as sbl,
		CAST((SELECT COUNT(pmf) FROM pos_forms WHERE pmf::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as pmf,
		CAST((SELECT COUNT(pmm) FROM pos_forms WHERE pmm::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as pmm,
		CAST((SELECT COUNT(ticket) FROM pos_forms WHERE ticket::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as ticket,
		CAST((SELECT COUNT(mtc) FROM pos_forms WHERE mtc::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as mtc,
		CAST((SELECT COUNT(ws) FROM pos_forms WHERE ws::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as ws,
		CAST((SELECT COUNT(mast) FROM pos_forms WHERE mast::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as mast,
		CAST((SELECT COUNT(oris) FROM pos_forms WHERE oris::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as oris,
		CAST((SELECT COUNT(elite) FROM pos_forms WHERE elite::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as elite,
		CAST((SELECT COUNT(yes) FROM pos_forms WHERE yes::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as yes,
		CAST((SELECT COUNT(time) FROM pos_forms WHERE time::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float as decimal(40,1)) * 100 as time
	FROM pos_forms
		INNER JOIN areas ON pos_forms.area_id=areas.id
		INNER JOIN provinces ON pos_forms.province_id=provinces.id
		WHERE "provinces"."name"=? AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP 
		AND ? ::TIMESTAMP 
		GROUP BY areas.name;
	`
	var chartData []models.NDChartData
	database.DB.Raw(sql1, province, start_date, end_date).Scan(&chartData)

	fmt.Println("chartData", chartData)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})

}

func PosByArea(c *fiber.Ctx) error {
	province := c.Params("province")
	area := c.Params("area")
	start_date := c.Params("start_date")
	end_date := c.Params("end_date")

	sql1 := ` SELECT
	(SELECT COUNT(equateur) FROM pos_forms WHERE equateur::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as eq, 
		(SELECT COUNT(dhl) FROM pos_forms WHERE dhl::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as dhl,
		(SELECT COUNT(ar) FROM pos_forms WHERE ar::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as ar,
		(SELECT COUNT(sbl) FROM pos_forms WHERE sbl::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as sbl,
		(SELECT COUNT(pmf) FROM pos_forms WHERE pmf::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as pmf,
		(SELECT COUNT(pmm) FROM pos_forms WHERE pmm::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as pmm,
		(SELECT COUNT(ticket) FROM pos_forms WHERE ticket::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as ticket,
		(SELECT COUNT(mtc) FROM pos_forms WHERE mtc::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as mtc,
		(SELECT COUNT(ws) FROM pos_forms WHERE ws::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as ws,
		(SELECT COUNT(mast) FROM pos_forms WHERE mast::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as mast,
		(SELECT COUNT(oris) FROM pos_forms WHERE oris::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as oris,
		(SELECT COUNT(elite) FROM pos_forms WHERE elite::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as elite,
		(SELECT COUNT(yes) FROM pos_forms WHERE yes::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as yes,
		(SELECT COUNT(time) FROM pos_forms WHERE time::float > 0)::float / (SELECT COUNT(pos_forms.id) FROM pos_forms)::float * 100 as time
	FROM pos_forms
		INNER JOIN areas ON pos_forms.area_id=areas.id
		INNER JOIN provinces ON pos_forms.province_id=provinces.id
		WHERE "provinces"."name"=? AND "areas"."name"=? AND "pos_forms"."created_at" BETWEEN ? ::TIMESTAMP 
		AND ? ::TIMESTAMP 
		GROUP BY areas.name;
	`
	var chartData models.NDChartDataBar
	database.DB.Raw(sql1, province, area, start_date, end_date).Scan(&chartData) 

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}
