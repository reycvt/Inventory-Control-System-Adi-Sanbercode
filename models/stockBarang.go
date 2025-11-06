package models

import "gorm.io/gorm"

type StockBarang struct {
	gorm.Model
	BarangID     uint
	StatusBarang bool         `gorm:"default:false"`
	Qty          int          `grom:"type:int(10)"`
	MasterBarang MasterBarang `gorm:"foreignKey:BarangID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
