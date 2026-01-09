package reviewrepo

import "fmt"

// GetBuyerIDByUserID - Get buyer_id from user_id
func (r *ReviewRepository) GetBuyerIDByUserID(userID uint) (uint, error) {
	var buyer struct {
		ID uint
	}
	err := r.db.Table("buyers").
		Select("id").
		Where("user_id = ?", userID).
		Where("deleted_at IS NULL").
		First(&buyer).Error

	if err != nil {
		return 0, fmt.Errorf("buyer account not found")
	}

	return buyer.ID, nil
}
