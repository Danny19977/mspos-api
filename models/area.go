package models

import "gorm.io/gorm"

type Area struct {
	gorm.Model

	Name       string `gorm:"not null" json:"name"`
	ProvinceID uint   `gorm:"foreignKey:province_id" json:"province_id"`
	SupID      uint   `gorm:"foreignKey:sup_id" json:"sup_id"`
	Signature  string `json:"signature"`
}

type AreaPaginate struct { 
	Id       uint `json:"id"`
	Name       string `json:"name"`
	Province string   `json:"province"`
	Sup      string   `json:"sup"`
	Signature  string `json:"signature"`
}

type AreaDropDown struct { 
	Id  uint `json:"id"`
	Name  string  `json:"name"`
	ProvinceID uint `json:"province_id"`
}


func (p *Area) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Area{}).Count(&total)
	return total
}

func (p *Area) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Area{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}
