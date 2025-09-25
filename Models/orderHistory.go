package models

import "time"

type OrderStatusEnum string

const (
	OrderPending    OrderStatusEnum = "pending"
	OrderProcessing OrderStatusEnum = "processing"
	OrderShipped    OrderStatusEnum = "shipped"
	OrderDelivered  OrderStatusEnum = "delivered"
	OrderCancelled  OrderStatusEnum = "cancelled"
)

type OrderHistory struct {
	ID        uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   uint            `json:"order_id" gorm:"not null;index:idx_order_id"`
	Status    OrderStatusEnum `json:"status" gorm:"type:varchar(20);not null;default:'pending';check:status IN ('pending','processing','shipped','delivered','cancelled')"`
	Comment   string          `json:"comment" gorm:"type:text"`
	UpdatedBy string          `json:"updated_by" gorm:"size:50"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`

	Order Order `json:"order" gorm:"foreignKey:OrderID"`
}
