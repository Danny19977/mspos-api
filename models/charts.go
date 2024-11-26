package models

type NDChartData struct {
    Area   string
	Eq     float64
	// Sold   []int
	Dhl    float64
	Ar     float64
	Sbl    float64
	Pmf    float64
	Pmm    float64
	Ticket float64
	Mtc    float64
	Ws     float64
	Mast   float64
	Oris   float64
	Elite  float64
	Ck     float64
	Yes    float64
	Time   float64
}

type NDAverage struct {
	Brand string
	Pourcent float64
}

type NdByYear struct {
	Month string 
	Eq float64  
}

type NDChartDataBar struct {
	Eq     float64 
	Dhl    float64
	Ar     float64
	Sbl    float64
	Pmf    float64
	Pmm    float64
	Ticket float64
	Mtc    float64
	Ws     float64
	Mast   float64
	Oris   float64
	Elite  float64
	Ck     float64
	Yes    float64
	Time   float64
}

type SummaryCount struct {
	Count int64
}

type SosPieChart struct {
	Province string
	Eq int64
}


type TrackingVisitDRSChart struct {
	Province string
	Nd int64
	Sos int64
	Oos int64
	Dr int64
	Visit int64
	Obj int64
	Perf int64
}

type SumChartBar struct {
	Province string
	Nd int64
	Sos int64
	Oos int64 
}

type SosPieChartArea struct {
	Area string
	Eq int64
}

type SummaryBetterDR struct {
	Fullname string
	Province string
	Area string
	Ventes int64
}

type StatusEquip struct {
	Equipement string
	Count int64
}

type GoogleMap struct {
	Latitude string
	Longitude string
	Name string
}

type PriceChart struct {
	Price string
	Count int64
}