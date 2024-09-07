package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)



func NdDashboard(c *fiber.Ctx) error {
	rows, err := database.DBSQL.Query(`
		SELECT SUM(equateur) AS equateur, SUM(sold) AS sold, SUM(dhl) AS dhl, SUM(ar) AS ar, SUM(sbl) AS sbl, SUM(pmt) AS pmt, 
		SUM(pmm) AS pmm, SUM(ticket) AS ticket, 
		SUM(mtc) AS mtc, SUM(ws) AS ws, SUM(mast) AS mast, SUM(oris) AS oris, SUM(elite) AS elite, SUM(ck) AS ck, 
		SUM(yes) AS yes, SUM(time) AS time, name
		FROM pos_forms 
		INNER JOIN areas ON pos_forms.area_id=areas.id
		GROUP BY name;
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var chartData models.NDChartData
	for rows.Next() {
		var name string
		var equateur int
		var sold int
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
		err := rows.Scan(&equateur, &sold, &dhl, &ar, &sbl, &pmt, &pmm, &ticket, &mtc, &ws,
			&mast, &oris, &elite, &ck, &yes, &time, &name)
		if err != nil {
			return err
		}
		chartData.Name = append(chartData.Name, name)
		chartData.Equateur = append(chartData.Equateur, equateur)
		chartData.Sold = append(chartData.Sold, sold)
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
