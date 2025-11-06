package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Username  string `gorm:"size:50;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HashPassword(Password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	return err == nil
}
