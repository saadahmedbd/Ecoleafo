package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserID     uint   `json:"user_id" gorm:"uniqueIndex;not null"`
	Role       string `json:"role" gorm:"size:20;not null;default:'admin'"` // admin, super_admin
	FullName   string `json:"full_name" gorm:"size:100;not null"`
	Phone      string `json:"phone" gorm:"size:20"`
	Department string `json:"department" gorm:"size:50"` //sales, operation, support etc

	//permisssion
	CanManageUsers    bool `json:"can_manage_users" gorm:"default:true"`
	CanManageProducts bool `json:"can_manage_products" gorm:"default:true"`
	CanManageOrders   bool `json:"can_manage_orders" gorm:"default:true"`
	CanViewReports    bool `json:"can_view_reports" gorm:"default:true"`
	CanManageSettings bool `json:"can_manage_settings" gorm:"default:false"`

	IsActive  bool           `json:"is_active" gorm:"default:true"`
	LastLogin *time.Time     `json:"last_login"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	RegUser RegUser `json:"reg_user" gorm:"foreignKey:UserID;references:ID"`
}
