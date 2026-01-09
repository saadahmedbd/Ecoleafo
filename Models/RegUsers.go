package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegUser struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName   string     `json:"first_name" gorm:"size:100;not null;default:'N/A'"`
	LastName    string     `json:"last_name" gorm:"size:100;not null;default:'N/A'"`
	Email       string     `json:"email" gorm:"uniqueIndex;size:100;not null"`
	Password    string     `json:"password" gorm:"size:255;not null"`
	Gender      string     `json:"gender"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Role        string     `json:"role" gorm:"size:20;not null"` // buyer, seller, admin
	Phone       string     `json:"phone" gorm:"size:20;not null;default:'N/A'"`
	Avatar      string     `json:"avatar" gorm:"size:255"`

	// Google OAuth fields
	GoogleID     string `json:"google_id" gorm:"size:255;index"`              // Store Google's user ID
	AuthProvider string `json:"auth_provider" gorm:"size:20;default:'local'"` // local, google

	IsActive         bool           `json:"is_active" gorm:"default:true"`       // status
	IsVerified       bool           `json:"is_verified" gorm:"default:false"`    //  verification
	EmailVerified    bool           `json:"email_verified" gorm:"default:false"` //  email verification
	ProfilePhoto     string         `json:"profile_photo" gorm:"size:255"`
	TwoFactorEnabled bool           `json:"two_factor_enabled" gorm:"default:false"` // 2FA status
	LastLoginAt      *time.Time     `json:"last_login_at"`                           //  last login tracking
	UpdatedAt        time.Time      `json:"updated_at"`                              // updated_at
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`                 //  soft delete
	Created_At       time.Time      `json:"created_at"`

	// Relationships
	SellerProfile *User  `json:"seller_profile,omitempty" gorm:"foreignKey:UserId;references:ID"`
	BuyerProfile  *Buyer `json:"buyer_profile,omitempty" gorm:"foreignKey:UserId;references:ID"`
	AdminProfile  *Admin `json:"admin_profile,omitempty" gorm:"foreignKey:UserID;references:ID"`
}

func (u *RegUser) CheckPassword(password string) error {
	// For Google OAuth users without password
	if u.Password == "" {
		return fmt.Errorf("this account uses Google sign-in")
	}
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
