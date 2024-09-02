package models


import "gorm.io/gorm"

type Manager struct {
	gorm.Model

	Name  string `gorm:"not null" json:"name"`
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