package database

import (
	"fmt"
	"inverntory-adi-sanbercode/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("adikanna123"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Gagal hash password: %v", err)
		}
		admin := models.User{

			Name:     "Adi Kannatasik",
			Username: "admin",
			Password: string(hashedPassword),
		}

		DB.Create(&admin)
		fmt.Println("Seeder Berhasil")

	} else {
		fmt.Println("Seeder Gagal")
	}
}
