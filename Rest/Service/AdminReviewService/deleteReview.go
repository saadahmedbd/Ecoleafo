package adminreviewservice

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
)

// Delete review
func (s *adminReviewService) DeleteReview(id uint, adminID uint) error {
	// Get review
	review, err := s.adminreviewRepo.GetReviewByID(id)
	if err != nil {
		return fmt.Errorf("review not found: %w", err)
	}

	if err := s.adminreviewRepo.DeleteReview(id); err != nil {
		return err
	}

	// Update product rating
	s.updateProductRating(review.ProductID)

	// Log action
	s.auditlogRepo.Create(&models.AuditLog{
		ActorID:     &adminID,
		Action:      "review_deleted",
		EntityType:  "review",
		EntityID:    &id,
		Description: fmt.Sprintf("Deleted review #%d for product #%d", id, review.ProductID),
	})

	return nil
}
