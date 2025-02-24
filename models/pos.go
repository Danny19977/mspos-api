package models

import "gorm.io/gorm"

type Pos struct {
	gorm.Model

	Name               string   `gorm:"not null" json:"name"` // Celui qui vend
	Shop               string   `json:"shop"`                 // Nom du shop
	Postype            string   `json:"postype"`              // Type de POS
	Manager            string   `json:"manager"`              // name of the onwer of the pos
	Commune            string   `json:"commune"`
	Avenue             string   `json:"avenue"`
	Quartier           string   `json:"quartier"`
	Reference          string   `json:"reference"`
	Telephone          string   `json:"telephone"`
	Eparasol           bool     `json:"eparasol"`
	Etable             bool     `json:"etable"`
	Ekiosk             bool     `json:"ekiosk"`
	InputGroupSelector string   `json:"inputgroupselector"`
	Cparasol           bool     `json:"cparasol"`
	Ctable             bool     `json:"ctable"`
	Ckiosk             bool     `json:"ckiosk"`
	ProvinceID         uint     `json:"province_id"`
	Province           Province `gorm:"foreignKey:ProvinceID"`
	AreaID             uint     `json:"area_id"`
	Area               Area     `gorm:"foreignKey:ProvinAreaIDceID"`
	UserID             uint     `json:"user_id"`
	User               User     `gorm:"foreignKey:UserID"`
	Status             bool     `json:"status"`
	Signature          string   `json:"signature"`

	Dr string `json:"dr"`
}

type PosPaginate struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`    // Celui qui vend
	Shop               string `json:"shop"`    // Nom du shop
	Manager            string `json:"manager"` // name of the onwer of the pos
	Commune            string `json:"commune"`
	Avenue             string `json:"avenue"`
	Quartier           string `json:"quartier"`
	Reference          string `json:"reference"`
	Telephone          string `json:"telephone"`
	Eparasol           bool   `json:"eparasol"`
	Etable             bool   `json:"etable"`
	Ekiosk             bool   `json:"ekiosk"`
	InputGroupSelector string `json:"inputgroupselector"`
	Cparasol           bool   `json:"cparasol"`
	Ctable             bool   `json:"ctable"`
	Ckiosk             bool   `json:"ckiosk"`
	Province           string `json:"province"`
	Area               string `json:"area"`
	Dr                 string `json:"dr"`
	Status             bool   `json:"status"`
	Signature          string `json:"signature"`
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
