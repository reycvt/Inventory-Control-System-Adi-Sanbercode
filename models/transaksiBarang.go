package models

import "gorm.io/gorm"

type TransaksiBarang struct {
	gorm.Model

	BarangID     uint
	Type         string       `gorm:"type:varchar(10);not null;check:type in ('IN','OUT')"`
	Qty          int          `gorm:"type:int(10)"`
	Keterangan   string       `gorm:"type:text"`
	MasterBarang MasterBarang `gorm:"foreignKey:BarangID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
