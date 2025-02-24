package models

import "gorm.io/gorm"

type Asm struct {
	gorm.Model

	Name       string   `gorm:"not null" json:"name"`
	CountryID  uint     `json:"country_id"`
	Country    Country  `gorm:"foreignKey:CountryID"`
	ProvinceID uint     `json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceID"`
	Signature  string   `json:"signature"`

	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID"`

	// Sups []Sup `gorm:"foreignKey:AmsID"` // One to many relationship
}

type AsmPaginate struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Province  string `json:"province"`
	Signature string `json:"signature"`
}

func (p *Asm) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Asm{}).Count(&total)
	return total
}

func (p *Asm) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Asm{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}
