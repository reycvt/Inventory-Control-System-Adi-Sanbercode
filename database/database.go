package database

import (
	"fmt"
	"inverntory-adi-sanbercode/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Tidak ada env")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
