package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegUser struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName  string    `json:"first_name" gorm:"size:100;not null;default:'N/A'"`
	LastName   string    `json:"last_name" gorm:"size:100;not null;default:'N/A'"`
	Email      string    `json:"email" gorm:"uniqueIndex;size:100;not null"`
	Password   string    `json:"password" gorm:"size:255;not null"`
	Role       string    `json:"role" gorm:"size:20;not null"` // buyer, seller, admin
	Created_At time.Time `json:"created_at`
}

func (u *RegUser) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
