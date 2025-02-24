package models

type Country struct {
	Name      string    `gorm:"not null" json:"name"`
	Signature string    `json:"signature"`
	Provinces []Province `gorm:"foreignKey:CountryID"`
	Users     []User    `gorm:"foreignKey:CountryID"`
	Areas     []Area    `gorm:"foreignKey:CountryID"`
	Sups      []Sup     `gorm:"foreignKey:CountryID"`
	Asms      []Asm     `gorm:"foreignKey:CountryID"`
	Posforms  []PosForm `gorm:"foreignKey:CountryID"`
}
