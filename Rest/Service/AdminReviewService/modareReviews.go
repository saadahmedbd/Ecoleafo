package adminreviewservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
)

// Moderate review (approve or reject)
func (s *adminReviewService) ModerateReview(id uint, req *adminreviewdto.ModerateReviewRequest, adminID uint) error {
	// Get review
	review, err := s.adminreviewRepo.GetReviewByID(id)
	if err != nil {
		return fmt.Errorf("review not found: %w", err)
	}

	if review.Status != "pending" && review.Status != "approved" {
		return errors.New("only pending or approved reviews can be moderated")
	}

	var moderationErr error
	if req.Action == "approve" {
		moderationErr = s.adminreviewRepo.ApproveReview(id, adminID)

		// Update product average rating
		if moderationErr == nil {
			s.updateProductRating(review.ProductID)
		}

		// Log action
		s.auditlogRepo.Create(&models.AuditLog{
			ActorID:     &adminID,
			Action:      "review_approved",
			EntityType:  "review",
			EntityID:    &id,
			Description: fmt.Sprintf("Approved review #%d for product #%d", id, review.ProductID),
		})

	} else if req.Action == "reject" {
		if req.Reason == "" {
			return errors.New("rejection reason is required")
		}

		moderationErr = s.adminreviewRepo.RejectReview(id, req.Reason, adminID)

		// Log action
		s.auditlogRepo.Create(&models.AuditLog{
			ActorID:     &adminID,
			Action:      "review_rejected",
			EntityType:  "review",
			EntityID:    &id,
			Description: fmt.Sprintf("Rejected review #%d: %s", id, req.Reason),
		})
	} else {
		return errors.New("invalid action: must be 'approve' or 'reject'")
	}

	return moderationErr
}
