package models

import "gorm.io/gorm"

type MasterBarang struct {
	gorm.Model
	NamaBarang string `gorm:"type:varchar(255);not null"`
	Type       string `gorm:"type:varchar(50)"`
	Keterangan string `gorm:"type:text"`
}
