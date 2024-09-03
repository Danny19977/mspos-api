package models

import "gorm.io/gorm"

type Province struct {
	gorm.Model

	Name  string  `json:"name"`
	// Users []User 
	// Areas []Area 
	// Sups []Sup 
	// Asms []Asm 
	// Posforms []PosForm 
	Signature    string `json:"signature"`
}

func (p *Province) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Province{}).Count(&total)
	return total
}

func (p *Province) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Province{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}