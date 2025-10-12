package orderrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *orderRepository) FindByOrderNumber(orderNumber string) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Buyer").
		Preload("OrderItems.Product").
		Preload("OrderItems.Seller").
		Where("order_number = ?", orderNumber).
		First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

	}
	return &order, nil
}
