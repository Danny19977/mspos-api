package models

import (
	"gorm.io/gorm"
)

type PosFormItems struct {
	gorm.Model

	PosFormID  uint     `json:"posform_id"`             // Foreign key (belongs to), tag `index` will create index for this column
	PosForm    PosForm  `gorm:"foreignKey:PosFormID"`    // POS Form of the POS
	BrandID    uint     `json:"brand_id"`                // Foreign key (belongs to), tag `index` will create index for this column
	Brand      Brand    `gorm:"foreignKey:BrandID"`      // Brand of the POS
	Number     uint     `gorm:"not null" json:"number"`  // Number of Frudes fond in a POS
	Counter    uint     `gorm:"not null" json:"counter"` // Counter for ND
	ProvinceID uint     `json:"province_id"`             // Foreign key (belongs to), tag `index` will create index for this column
	Province   Province `gorm:"foreignKey:ProvinceID"`   // Province of the POS
	AreaID     uint     `json:"area_id"`                 // Foreign key (belongs to), tag `index` will create index for this column
	Area       Area     `gorm:"foreignKey:AreaID"`       // Area of the POS
	SubArea    SubArea  `gorm:"foreignKey:SubAreaID"`    // Subarea of the POS
	SubAreaID  uint     `json:"subArea_id"`              // Foreign key (belongs to), tag `index` will create index for this column
}
