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

	sql := fmt.Sprintf(`
		SELECT areas.name AS area,
		cast((SELECT COUNT(equateur) FROM pos_forms WHERE equateur > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as eq, 
		cast((SELECT COUNT(dhl) FROM pos_forms WHERE dhl > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as dhl,
		cast((SELECT COUNT(ar) FROM pos_forms WHERE ar > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as ar,
		cast((SELECT COUNT(sbl) FROM pos_forms WHERE sbl > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as sbl,
		cast((SELECT COUNT(pmt) FROM pos_forms WHERE pmt > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as pmt,
		cast((SELECT COUNT(pmm) FROM pos_forms WHERE pmm > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as pmm,
		cast((SELECT COUNT(ticket) FROM pos_forms WHERE ticket > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as ticket,
		cast((SELECT COUNT(mtc) FROM pos_forms WHERE mtc > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as mtc,
		cast((SELECT COUNT(ws) FROM pos_forms WHERE ws > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as ws,
		cast((SELECT COUNT(mast) FROM pos_forms WHERE mast > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as mast,
		cast((SELECT COUNT(oris) FROM pos_forms WHERE oris > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as oris,
		cast((SELECT COUNT(elite) FROM pos_forms WHERE elite > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as elite,
		cast((SELECT COUNT(ck) FROM pos_forms WHERE ck > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as ck,
		cast((SELECT COUNT(yes) FROM pos_forms WHERE yes > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as yes,
		cast((SELECT COUNT(time) FROM pos_forms WHERE time > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as time
		FROM pos_forms
		INNER JOIN areas ON pos_forms.area_id=areas.id
		INNER JOIN provinces ON pos_forms.province_id=provinces.id
		WHERE "provinces"."name"='%s' && created_at BETWEEN %s' ::TIMESTAMP AND %s' ::TIMESTAMP 
		GROUP BY areas.name;
	`, province, start_date, end_date)

	rows, err := database.DBSQL.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var chartData models.NDChartData
	for rows.Next() {
		var area string
		var eq int
		// var sold int
		var dhl int
		var ar int
		var sbl int
		var pmt int
		var pmm int
		var ticket int
		var mtc int
		var ws int
		var mast int
		var oris int
		var elite int
		var ck int
		var yes int
		var time int
		err := rows.Scan(&area, &eq, &dhl, &ar, &sbl, &pmt, &pmm, &ticket, &mtc, &ws,
			&mast, &oris, &elite, &ck, &yes, &time)
		if err != nil {
			return err
		}
		chartData.Area = append(chartData.Area, area)
		chartData.Eq = append(chartData.Eq, eq)
		chartData.Dhl = append(chartData.Dhl, dhl)
		chartData.Ar = append(chartData.Ar, ar)
		chartData.Sbl = append(chartData.Sbl, sbl)
		chartData.Pmt = append(chartData.Pmt, pmt)
		chartData.Pmm = append(chartData.Pmm, pmm)
		chartData.Ticket = append(chartData.Ticket, ticket)
		chartData.Mtc = append(chartData.Mtc, mtc)
		chartData.Ws = append(chartData.Ws, ws)
		chartData.Mast = append(chartData.Mast, mast)
		chartData.Oris = append(chartData.Oris, oris)
		chartData.Elite = append(chartData.Elite, elite)
		chartData.Ck = append(chartData.Ck, ck)
		chartData.Yes = append(chartData.Yes, yes)
		chartData.Time = append(chartData.Time, time)
	}

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

	sql := fmt.Sprintf(`
		SELECT name AS area,
		cast((SELECT COUNT(equateur) FROM pos_forms WHERE equateur > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as eq, 
		cast((SELECT COUNT(dhl) FROM pos_forms WHERE dhl > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as dhl,
		cast((SELECT COUNT(ar) FROM pos_forms WHERE ar > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as ar,
		cast((SELECT COUNT(sbl) FROM pos_forms WHERE sbl > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as sbl,
		cast((SELECT COUNT(pmt) FROM pos_forms WHERE pmt > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as pmt,
		cast((SELECT COUNT(pmm) FROM pos_forms WHERE pmm > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as pmm,
		cast((SELECT COUNT(ticket) FROM pos_forms WHERE ticket > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as ticket,
		cast((SELECT COUNT(mtc) FROM pos_forms WHERE mtc > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as mtc,
		cast((SELECT COUNT(ws) FROM pos_forms WHERE ws > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as ws,
		cast((SELECT COUNT(mast) FROM pos_forms WHERE mast > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as mast,
		cast((SELECT COUNT(oris) FROM pos_forms WHERE oris > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as oris,
		cast((SELECT COUNT(elite) FROM pos_forms WHERE elite > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as elite,
		cast((SELECT COUNT(ck) FROM pos_forms WHERE ck > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as ck,
		cast((SELECT COUNT(yes) FROM pos_forms WHERE yes > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as yes,
		cast((SELECT COUNT(time) FROM pos_forms WHERE time > 0)::float / COUNT(pos_forms.id::float)* 100 as decimal(40,0)) as time
		FROM pos_forms
		INNER JOIN areas ON pos_forms.area_id=areas.id
		INNER JOIN provinces ON pos_forms.province_id=provinces.id
		WHERE "provinces"."name"='%s' && "areas"."name"='%s' && created_at BETWEEN '%s' ::TIMESTAMP AND '%s' ::TIMESTAMP 
		GROUP BY "areas"."name";
	`, province, area, start_date, end_date)

	rows, err := database.DBSQL.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var chartData models.NDChartData
	for rows.Next() {
		var area string
		var eq int
		// var sold int
		var dhl int
		var ar int
		var sbl int
		var pmt int
		var pmm int
		var ticket int
		var mtc int
		var ws int
		var mast int
		var oris int
		var elite int
		var ck int
		var yes int
		var time int
		err := rows.Scan(&area, &eq, &dhl, &ar, &sbl, &pmt, &pmm, &ticket, &mtc, &ws,
			&mast, &oris, &elite, &ck, &yes, &time)
		if err != nil {
			return err
		}

		chartData.Area = append(chartData.Area, area)
		chartData.Eq = append(chartData.Eq, eq)
		chartData.Dhl = append(chartData.Dhl, dhl)
		chartData.Ar = append(chartData.Ar, ar)
		chartData.Sbl = append(chartData.Sbl, sbl)
		chartData.Pmt = append(chartData.Pmt, pmt)
		chartData.Pmm = append(chartData.Pmm, pmm)
		chartData.Ticket = append(chartData.Ticket, ticket)
		chartData.Mtc = append(chartData.Mtc, mtc)
		chartData.Ws = append(chartData.Ws, ws)
		chartData.Mast = append(chartData.Mast, mast)
		chartData.Oris = append(chartData.Oris, oris)
		chartData.Elite = append(chartData.Elite, elite)
		chartData.Ck = append(chartData.Ck, ck)
		chartData.Yes = append(chartData.Yes, yes)
		chartData.Time = append(chartData.Time, time)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "chartData data",
		"data":    chartData,
	})
}
