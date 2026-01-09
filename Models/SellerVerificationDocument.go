package models

import (
	"time"

	"gorm.io/gorm"
)

type SellerVerificationDocument struct {
	ID              uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID        uint       `json:"seller_id" gorm:"not null;index"`
	DocumentType    string     `json:"document_type" gorm:"size:50;not null"` // identity, tax, bank
	DocumentURL     string     `json:"document_url" gorm:"size:500;not null"`
	Status          string     `json:"status" gorm:"size:50;default:'pending'"` // pending, verified, rejected
	RejectedAt      *time.Time `json:"rejected_at"`
	VerifiedAt      *time.Time `json:"verified_at"`
	RejectionReason string     `json:"rejection_reason" gorm:"type:text"`
	VerifiedBy      uint       `json:"verified_by"` // Admin user ID

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Seller User `json:"seller" gorm:"foreignKey:SellerID;references:ID"`
}

func (SellerVerificationDocument) TableName() string {
	return "seller_verification_documents"
}
