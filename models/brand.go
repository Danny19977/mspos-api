package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model

	Name       string   `gorm:"not null" json:"name"`
	ProvinceID uint     `json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceID"`
	Signature  string   `json:"signature"`

	PosFormItems []PosFormItems `gorm:"foreignKey:BrandID"`
}
