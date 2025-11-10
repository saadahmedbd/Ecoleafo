package reviewrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
)

// GetRatingBreakdown - Get rating breakdown for a product
func (r *ReviewRepository) GetRatingBreakdown(productID uint) (reviewdto.RatingBreakdown, error) {
	var breakdown reviewdto.RatingBreakdown

	rows, err := r.db.Model(&models.Review{}).
		Select("rating, COUNT(*) as count").
		Where("product_id = ?", productID).
		Group("rating").
		Rows()

	if err != nil {
		return breakdown, err
	}
	defer rows.Close()

	for rows.Next() {
		var rating, count int
		if err := rows.Scan(&rating, &count); err != nil {
			return breakdown, err
		}

		switch rating {
		case 5:
			breakdown.FiveStar = count
		case 4:
			breakdown.FourStar = count
		case 3:
			breakdown.ThreeStar = count
		case 2:
			breakdown.TwoStar = count
		case 1:
			breakdown.OneStar = count
		}
	}

	return breakdown, nil
}
