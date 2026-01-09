package orderrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *orderRepository) GetOrderItemByID(itemID uint) (*models.OrderItem, error) {
	var item models.OrderItem
	err := r.db.Preload("Product").Preload("Seller").Preload("Order").First(&item, itemID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order item not found")
		}
		return nil, err
	}
	return &item, nil
}
