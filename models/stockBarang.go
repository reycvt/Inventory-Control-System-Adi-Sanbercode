package models

import "gorm.io/gorm"

type StockBarang struct {
	gorm.Model
	BarangID     uint
	StatusBarang bool         `gorm:"default:false"`
	Qty          int          `gorm:"type:int(10)"`
	MasterBarang MasterBarang `gorm:"foreignKey:BarangID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
