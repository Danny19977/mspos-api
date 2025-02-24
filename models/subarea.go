package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type SubArea struct {
	gorm.Model

	Name       string         `gorm:"not null" json:"name"`
	Commune    pq.StringArray `gorm:"type:text[];null" json:"commune"`
	ProvinceID uint           `json:"province_id"`
	Province   Province       `gorm:"foreignKey:ProvinceID"`
	Signature  string         `json:"signature"`

	Sups []Sup `gorm:"foreignKey:ProvinceID"`
	Dr   []Dr  `gorm:"foreignKey:ProvinceID"`
	User []User `gorm:"foreignKey:SubAreaID"`
}
