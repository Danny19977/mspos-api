package models

import "gorm.io/gorm"

type RoutePlan struct {
	gorm.Model

	PosListID  []string `gorm:"type:text[];null" json:"pos_list_id"`
	UserID     uint     `json:"user_id"`
	User       User     `gorm:"foreignKey:UserID"`
	ProvinceID uint     `json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceID"`
	AreaID     uint     `json:"area_id"`
	Area       Area     `gorm:"foreignKey:AreaID"`
	SubAreaID  uint     `json:"sub_area_id"`
	SubArea    SubArea  `gorm:"foreignKey:SubAreaID"`
	TotalPOS   int      `json:"total_pos"`
	Signature  string   `json:"signature"`
}
