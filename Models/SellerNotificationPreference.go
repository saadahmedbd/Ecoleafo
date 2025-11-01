package models

import (
	"time"

	"gorm.io/gorm"
)

type SellerNotificationPreference struct {
	ID       uint `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID uint `json:"seller_id" gorm:"uniqueIndex;not null"` // References User.ID

	// Order Notifications
	OrderEmail bool `json:"order_email" gorm:"default:true"`
	OrderSMS   bool `json:"order_sms" gorm:"default:false"`
	OrderPush  bool `json:"order_push" gorm:"default:true"`

	// Message Notifications
	MessageEmail bool `json:"message_email" gorm:"default:true"`
	MessageSMS   bool `json:"message_sms" gorm:"default:false"`
	MessagePush  bool `json:"message_push" gorm:"default:true"`

	// Marketing Notifications
	MarketingEmail bool `json:"marketing_email" gorm:"default:false"`
	MarketingSMS   bool `json:"marketing_sms" gorm:"default:false"`
	MarketingPush  bool `json:"marketing_push" gorm:"default:false"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Seller User `json:"seller" gorm:"foreignKey:SellerID;references:ID"`
}

func (SellerNotificationPreference) TableName() string {
	return "seller_notification_preferences"
}
