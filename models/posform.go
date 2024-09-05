package models

import "gorm.io/gorm"

type PosForm struct {
	gorm.Model

	IdUnique string `json:"id_unique"`
	Equateur int64  `json:"equateur"`
	Sold     int64  `json:"sold"`
	Dhl      int64  `json:"dhl"`
	Ar       int64  `json:"ar"`
	Sbl      int64  `json:"sbl"`
	Pmt      int64  `json:"pmt"`
	Pmm      int64  `json:"pmm"`
	Ticket   int64  `json:"ticket"`
	Mtc      int64  `json:"mtc"`
	Ws       int64  `json:"ws"`
	Mast     int64  `json:"mast"`
	Oris     int64  `json:"oris"`
	Elite    int64  `json:"elite"`
	Ck       int64  `json:"ck"`
	Yes      int64  `json:"yes"`
	Time     int64  `json:"time"`
	Comment  string `json:"comment"`

	ProvinceID uint `gorm:"foreignKey:province_id" json:"province_id"`
	UserID     uint `gorm:"foreignKey:user_id" json:"user_id"`
	AreaID     uint `gorm:"foreignKey:area_id" json:"area_id"`
	SupID      uint `gorm:"foreignKey:sup_id" json:"sup_id"`
	PosID      uint `gorm:"foreignKey:pos_id" json:"pos_id"`

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
