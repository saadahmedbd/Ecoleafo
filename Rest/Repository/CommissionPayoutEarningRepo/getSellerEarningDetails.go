package commissionpayoutearningrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *commissionRepository) GetSellerEarningDetails(sellerID uint) (*models.SellerEarningsSummary, []models.Order, error) {
	var earnings models.SellerEarningsSummary
	err := r.db.Preload("Seller.RegUser").Where("seller_id = ?", sellerID).First(&earnings).Error
	if err != nil {
		return nil, nil, err
	}

	var recentOrders []models.Order
	err = r.db.Where("seller_id = ?", sellerID).
		Order("created_at DESC").
		Limit(10).
		Find(&recentOrders).Error

	return &earnings, recentOrders, err
}
