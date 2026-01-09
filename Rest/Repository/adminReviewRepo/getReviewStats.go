package adminreviewrepo

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
)

// Get review statistics
func (r *adminReviewRepository) GetReviewStats() (map[string]interface{}, error) {
	var total, pending, approved, rejected, reported int64
	var avgRating float64

	r.db.Model(&models.Review{}).Count(&total)
	r.db.Model(&models.Review{}).Where("status = ?", "pending").Count(&pending)
	r.db.Model(&models.Review{}).Where("status = ?", "approved").Count(&approved)
	r.db.Model(&models.Review{}).Where("status = ?", "rejected").Count(&rejected)
	r.db.Model(&models.Review{}).Where("is_reported = ?", true).Count(&reported)
	r.db.Model(&models.Review{}).Select("AVG(rating)").Where("status = ?", "approved").Scan(&avgRating)

	// Rating distribution
	var ratingDist []struct {
		Rating int
		Count  int64
	}
	r.db.Model(&models.Review{}).
		Select("rating, COUNT(*) as count").
		Where("status = ?", "approved").
		Group("rating").
		Scan(&ratingDist)

	distribution := make(map[string]int)
	for _, rd := range ratingDist {
		distribution[fmt.Sprintf("%d_star", rd.Rating)] = int(rd.Count)
	}

	return map[string]interface{}{
		"total":               total,
		"pending":             pending,
		"approved":            approved,
		"rejected":            rejected,
		"reported":            reported,
		"average_rating":      avgRating,
		"rating_distribution": distribution,
	}, nil
}
