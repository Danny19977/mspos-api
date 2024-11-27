package models

import (
	"time"

	"gorm.io/gorm"
)

type UserLogs struct {
	gorm.Model

	Name        string `gorm:"not null" json:"name"`
	UserID      uint   `gorm:"foreignKey:user_id" json:"user_id"`
	Action      string `gorm:"not null" json:"action"`
	Description string `gorm:"not null" json:"description"`
	Signature   string `json:"signature"`

	Fullname    string    `json:"fullname"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserLogPaginate struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      uint      `json:"user_id"`
	Fullname    string    `json:"fullname"`
	Title       string    `json:"title"`
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
