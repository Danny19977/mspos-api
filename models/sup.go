package models

import "gorm.io/gorm"

type Sup struct {
	gorm.Model

	CountryID  uint     `json:"country_id"`
	Country    Country  `gorm:"foreignKey:CountryID"`
	ProvinceID uint     `json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceID"`
	AsmID      uint     `json:"asm_id"`
	Asms       Asm      `gorm:"foreignKey:AsmID"`
	Signature  string   `json:"signature"`
	AreaID     uint     `json:"area_id"`
	Area       Area     `gorm:"foreignKey:AreaID"`

	UserID uint `json:"user_id"` 
	User   User `gorm:"foreignKey:UserID"`
	
	Dr      []Dr      `gorm:"foreignKey:SupID"`
	Cyclo   []Cyclo   `gorm:"foreignKey:SupID"`
	Subarea []SubArea `gorm:"foreignKey:SupID"`
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
