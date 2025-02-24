package models


import "gorm.io/gorm"

type Manager struct {
	gorm.Model

	Title     string  `json:"title"` // Example Head of Sales, Support, Manager, etc
	CountryID uint    `json:"country_id"`
	Country   Country `gorm:"foreignKey:CountryID"`
	UserID    uint    `json:"user_id"`
	User      User    `gorm:"foreignKey:HosID"`
	Signature string  `json:"signature"`
}

type ManagerPaginate struct { 
	Id  uint `json:"id"`
	Name  string `json:"name"`
	Signature    string `json:"signature"`
}

func (p *Manager) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Manager{}).Count(&total)
	return total
}

func (p *Manager) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Manager{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}