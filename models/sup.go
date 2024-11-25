package models

import "gorm.io/gorm"

type Sup struct {
	gorm.Model

	Name       string `gorm:"not null" json:"name"`
	ProvinceID uint   `gorm:"not null" json:"province_id"`
	AsmID      uint   `json:"asm_id"`
	Signature  string `json:"signature"`
	Province  string `json:"province"`
	Asm       string `json:"asm"`
}

type SupPaginate struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Province  string `json:"province"`
	Asm       string `json:"asm"`
	Signature string `json:"signature"`
}

func (p *Sup) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Sup{}).Count(&total)
	return total
}

func (p *Sup) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Sup{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}
