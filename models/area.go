package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Area struct {
	gorm.Model

	Name       string         `gorm:"not null" json:"name"`
	Commune    pq.StringArray `gorm:"type:text[];null" json:"commune"`
	ProvinceID uint           `gorm:"foreignKey:province_id" json:"province_id"`
	SupID      uint           `gorm:"foreignKey:sup_id" json:"sup_id"`
	Signature  string         `json:"signature"`

	Province  string `json:"province"`
	Sup       string `json:"sup"` 
}

type AreaPaginate struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Province  string `json:"province"`
	Commune   string `json:"commune"`
	Sup       string `json:"sup"`
	Signature string `json:"signature"`
}

type AreaDropDown struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Commune    string `json:"commune"`
	ProvinceID uint   `json:"province_id"`
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
