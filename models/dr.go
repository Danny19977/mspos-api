package models

import "gorm.io/gorm"

type Dr struct {
	gorm.Model

	CountryID  uint     `json:"country_id"`
	Country    Country  `gorm:"foreignKey:CountryID"`
	ProvinceID uint     `json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceID"`
	AmsID      uint     `json:"ams_id"`
	Asms       Asm      `gorm:"foreignKey:AmsID"`
	SupID      uint     `json:"sup_id"`
	Sup        Sup      `gorm:"foreignKey:SupID"`
	AreaID     uint     `json:"area_id"`
	Area       Area     `gorm:"foreignKey:AreaID"`
	Signature  string   `json:"signature"`

	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID"`

	Cyclo   []Cyclo   `gorm:"foreignKey:DrID"`
	SubArea []SubArea `gorm:"foreignKey:DrID"`
}
