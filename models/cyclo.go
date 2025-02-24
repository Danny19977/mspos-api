package models

import "gorm.io/gorm"

type Cyclo struct {
	gorm.Model

	CountryID  uint     `json:"country_id"`
	Country    Country  `gorm:"foreignKey:CountryID"`
	ProvinceID uint     `json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceID"`
	AmsID      uint     `json:"ams_id"`
	Asms       Asm      `gorm:"foreignKey:AmsID"`
	SupID      uint     `json:"sup_id"`
	Sups       Sup      `gorm:"foreignKey:SupID"`
	DrID       uint     `json:"dr_id"`
	Dr         Dr       `gorm:"foreignKey:DrID"`
	AreaID     uint     `json:"area_id"`
	Area       Area     `gorm:"foreignKey:AreaID"`
	SubareaID  uint     `json:"SubArea_id"`
	SubArea    SubArea  `gorm:"foreignKey:SubAreaID"`
	Signature  string   `json:"signature"`

	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID"`
}
