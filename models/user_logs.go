package models

import "gorm.io/gorm"

type UserLogs struct {
	gorm.Model

	Name        string `gorm:"not null" json:"name"`
	UserID      uint
	Action      string `gorm:"not null" json:"action"`
	Description string `gorm:"not null" json:"description"`
	Signature    string `json:"signature"`
}

func (p *UserLogs) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&UserLogs{}).Count(&total)
	return total
}

func (p *UserLogs) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []UserLogs{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}