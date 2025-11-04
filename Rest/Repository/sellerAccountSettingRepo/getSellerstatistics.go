package selleraccountsettingrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	"gorm.io/gorm"
)

// GetSellerStatistics retrieves seller statistics
func (r *SellerAccountSettingRepository) GetSellerStatisticsStruct(sellerID uint) (*selleraccountsetting.SellerStatisticsResponse, error) {
	var seller models.User
	err := r.db.Preload("RegUser").Where("user_id = ?", sellerID).First(&seller).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}

	var activeProducts, totalReviews, pendingOrders, completedOrders int64

	//  Count active products
	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND is_active = ?", sellerID, true).
		Count(&activeProducts)

	//  Count pending orders (through joins)
	r.db.Table("orders").
		Joins("JOIN order_items ON order_items.order_id = orders.id").
		Joins("JOIN products ON order_items.product_id = products.id").
		Where("products.seller_id = ? AND orders.status = ?", sellerID, "pending").
		Group("orders.id").
		Count(&pendingOrders)

	//  Count completed orders (same logic)
	r.db.Table("orders").
		Joins("JOIN order_items ON order_items.order_id = orders.id").
		Joins("JOIN products ON order_items.product_id = products.id").
		Where("products.seller_id = ? AND orders.status = ?", sellerID, "completed").
		Group("orders.id").
		Count(&completedOrders)

	//  Count total reviews by seller’s products
	r.db.Table("reviews").
		Joins("JOIN products ON reviews.product_id = products.id").
		Where("products.seller_id = ?", sellerID).
		Count(&totalReviews)

	return &selleraccountsetting.SellerStatisticsResponse{
		TotalSales:      seller.TotalSales,
		TotalOrders:     seller.TotalOrders,
		ActiveProducts:  int(activeProducts),
		Rating:          seller.AverageRating,
		TotalReviews:    int(totalReviews),
		PendingOrders:   int(pendingOrders),
		CompletedOrders: int(completedOrders),
		TotalEarnings:   seller.TotalEarnings,
	}, nil
}
