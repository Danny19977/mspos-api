package models

import "gorm.io/gorm"

type Area struct {
	gorm.Model

	Shop               string `json:"shop"`
	Manager            string `json:"manager"` // name of the onwer of the pos
	Commune            string `json:"commune"`
	Avenue             string `json:"avenue"`
	Quartier           string `json:"quartier"`
	Reference          string `json:"reference"`
	Number             int64  `json:"number"`
	Eparasol           string `json:"eparasol"`
	Etable             string `json:"etable"`
	Ekiosk             bool   `json:"ekiosk"`
	InputGroupSelector string `json:"inputgroupselector"`
	Cparasol           string `json:"cparasol"`
	Ctable             string `json:"ctable"`
	Ckiosk             bool   `json:"Ckiosk"`
	ProvinceID         uint   `gorm:"foreignKey:province_id" json:"province_id"`
	UserID             uint   `gorm:"foreignKey:user_id" json:"user_id"`
	Signature          string `json:"signature"`
}

func (p *Area) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Area{}).Count(&total)
	return total
}

func (p *Area) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Area{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}
