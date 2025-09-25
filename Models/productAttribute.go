package models

type ProductAttribute struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID uint   `json:"product_id" gorm:"not null"`
	Name      string `json:"name" gorm:"size:100;not null"`
	Value     string `json:"value" gorm:"size:255;not null"`

	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}
