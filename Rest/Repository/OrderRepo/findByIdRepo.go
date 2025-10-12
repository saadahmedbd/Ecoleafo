package orderrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *orderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Buyer").
		Preload("OrderItems.Product").
		Preload("OrderItems.Seller").
		First(&order, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}
