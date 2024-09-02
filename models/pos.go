package models


import "gorm.io/gorm"

type Pos struct {
	gorm.Model

	Name  string `gorm:"not null" json:"name"`
	

	// Users []User  
	// Posforms []PosForm  
	Signature    string `json:"signature"`
}

func (p *Pos) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Pos{}).Count(&total)
	return total
}

func (p *Pos) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Pos{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}