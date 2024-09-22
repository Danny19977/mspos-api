package models

import "gorm.io/gorm"

type PosForm struct {
	gorm.Model

	IdUnique string `json:"id_unique"`
	Eq       int64  `json:"eq"`
	Eq1      int64  `gorm:"default: 0" json:"eq1"`
	Sold     int64  `json:"sold"`
	Dhl      int64  `json:"dhl"`
	Dhl1     int64  `gorm:"default: 0" json:"dhl1"`
	Ar       int64  `json:"ar"`
	Ar1      int64  `gorm:"default: 0" json:"ar1"`
	Sbl      int64  `json:"sbl"`
	Sbl1     int64  `gorm:"default: 0" json:"sbl1"`
	Pmf      int64  `json:"pmf"`
	Pmf1     int64  `gorm:"default: 0" json:"pmf1"`
	Pmm      int64  `json:"pmm"`
	Pmm1     int64  `gorm:"default: 0" json:"pmm1"`
	Ticket   int64  `json:"ticket"`
	Ticket1  int64  `gorm:"default: 0" json:"ticket1"`
	Mtc      int64  `json:"mtc"`
	Mtc1     int64  `gorm:"default: 0" json:"mtc1"`
	Ws       int64  `json:"ws"`
	Ws1      int64  `gorm:"default: 0" json:"ws1"`
	Mast     int64  `json:"mast"`
	Mast1    int64  `gorm:"default: 0" json:"mast1"`
	Oris     int64  `json:"oris"`
	Oris1    int64  `gorm:"default: 0" json:"oris1"`
	Elite    int64  `json:"elite"`
	Elite1   int64  `gorm:"default: 0" json:"elite1"`
	Yes      int64  `json:"yes"`
	Yes1     int64  `gorm:"default: 0" json:"yes1"`
	Time     int64  `json:"time"`
	Time1    int64  `gorm:"default: 0" json:"time1"`
	Comment  string `json:"comment"`

	ProvinceID uint `gorm:"foreignKey:province_id" json:"province_id"`
	UserID     uint `gorm:"foreignKey:user_id" json:"user_id"`
	AreaID     uint `gorm:"foreignKey:area_id" json:"area_id"`
	SupID      uint `gorm:"foreignKey:sup_id" json:"sup_id"`
	PosID      uint `gorm:"foreignKey:pos_id" json:"pos_id"`

	Signature string `json:"signature"`
}

type PosFormPaginate struct {
	Id  uint `json:"id"`
	IdUnique string `json:"id_unique"`
	Eq       int64  `json:"eq"`
	Sold     int64  `json:"sold"`
	Dhl      int64  `json:"dhl"`
	Ar       int64  `json:"ar"`
	Sbl      int64  `json:"sbl"`
	Pmf      int64  `json:"pmf"`
	Pmm      int64  `json:"pmm"`
	Ticket   int64  `json:"ticket"`
	Mtc      int64  `json:"mtc"`
	Ws       int64  `json:"ws"`
	Mast     int64  `json:"mast"`
	Oris     int64  `json:"oris"`
	Elite    int64  `json:"elite"`
	Yes      int64  `json:"yes"`
	Time     int64  `json:"time"`
	Comment  string `json:"comment"`

	Province string `json:"province"`
	User     string `json:"user"`
	Area     string `json:"area"`
	Sup      string `json:"sup"`
	Pos      string `json:"pos"`

	Signature string `json:"signature"`
}


func (p *PosForm) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&PosForm{}).Count(&total)
	return total
}

func (p *PosForm) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []PosForm{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}
