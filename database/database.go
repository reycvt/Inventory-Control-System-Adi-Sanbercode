package database

import (
	"fmt"
	"inverntory-adi-sanbercode/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	host := "localhost"
	port := 5432
	user := "projectgin"
	password := "testgin123"
	dbname := "inventory-adi"

	dsn := fmt.Sprintf("host =%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal Koneksi ke database: %v", err)
	}
	DB = db
	fmt.Println("Berhasil konek ke PosgreSql")

}
func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.MasterBarang{},
		&models.StockBarang{},
		&models.TransaksiBarang{})
	if err != nil {
		log.Fatalf("Gagal melakykan migrasi: %v", err)
	}
	fmt.Println("Migrasi database berhasil")

}
