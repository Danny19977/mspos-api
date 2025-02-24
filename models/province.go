package models

import "gorm.io/gorm"

type Province struct {
	gorm.Model

	Name      string  `json:"name"`
	CountryID uint    `json:"country_id"`
	Country   Country `gorm:"foreignKey:CountryID"`
	Signature string  `json:"signature"`

	Brands   []Brand   `gorm:"foreignKey:ProvinceID"`
	Users    []User    `gorm:"foreignKey:ProvinceID"`
	Areas    []Area    `gorm:"foreignKey:ProvinceID"`
	SubArea  []SubArea `gorm:"foreignKey:ProvinceID"`
	Asms     []Asm     `gorm:"foreignKey:ProvinceID"`
	Sups     []Sup     `gorm:"foreignKey:ProvinceID"`
	Dr       []Dr      `gorm:"foreignKey:ProvinceID"`
	Cyclo    []Cyclo   `gorm:"foreignKey:ProvinceID"`
	Posforms []PosForm `gorm:"foreignKey:ProvinceID"`
}

type ProvincePaginate struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Signature string `json:"signature"`
}

type ProvinceDropDown struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (p *Province) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Province{}).Count(&total)
	return total
}

func (p *Province) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Province{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}
