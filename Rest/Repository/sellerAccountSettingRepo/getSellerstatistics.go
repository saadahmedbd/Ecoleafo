package selleraccountsettingrepo

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// GetSellerStatistics retrieves seller statistics
func (r *SellerAccountSettingRepository) GetSellerStatisticsStruct(sellerID uint) (*selleraccountsetting.SellerStatisticsResponse, error) {
	var seller models.User
	if err := r.db.First(&seller, sellerID).Error; err != nil {
		return nil, fmt.Errorf("seller not found: %w", err)
	}

	// Count queries
	var activeProducts, totalReviews, pendingOrders, completedOrders int64

	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND is_active = ?", sellerID, true).
		Count(&activeProducts)

	r.db.Table("orders").
		Where("seller_id = ? AND status = ?", sellerID, "pending").
		Count(&pendingOrders)

	r.db.Table("orders").
		Where("seller_id = ? AND status = ?", sellerID, "completed").
		Count(&completedOrders)

	r.db.Table("reviews").
		Joins("JOIN products ON reviews.product_id = products.id").
		Where("products.seller_id = ?", sellerID).
		Count(&totalReviews)

	// ✅ BEST: Return struct directly, no type assertions!
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
