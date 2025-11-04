package adminmangementservice

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	adminmangement "github.com/saadahmedbd/Treestore/Rest/Repository/AdminMangement"
	"github.com/saadahmedbd/Treestore/constants"
)

type SellerService struct {
	sellerRepo      *adminmangement.SellerRepository
	activityLogRepo *adminmangement.AuditLogRepository
}

func NewSellerService(sellerRepo *adminmangement.SellerRepository, activityLogRepo *adminmangement.AuditLogRepository) *SellerService {
	return &SellerService{
		sellerRepo:      sellerRepo,
		activityLogRepo: activityLogRepo,
	}
}

// GetAllSellers retrieves all sellers with pagination
func (s *SellerService) GetAllSellers(page, limit int, status string) ([]models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.sellerRepo.GetAll(page, limit, status)
}

// GetSellerByID retrieves a seller by ID
func (s *SellerService) GetSellerByID(id uint) (*models.User, error) {
	return s.sellerRepo.GetByID(id)
}

// GetPendingApprovals retrieves sellers pending approval
func (s *SellerService) GetPendingApprovals(page, limit int) ([]models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.sellerRepo.GetPendingApprovals(page, limit)
}

// ApproveSeller approves a seller registration
func (s *SellerService) ApproveSeller(id uint, adminID uint) error {
	seller, err := s.sellerRepo.GetByID(id)
	if err != nil {
		return err
	}

	if seller.ApprovalStatus == "approved" {
		return errors.New("seller is already approved")
	}

	err = s.sellerRepo.ApproveSeller(id, adminID)
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivityApprove, "seller", &id, "Approved seller registration")

	return nil
}

// RejectSeller rejects a seller registration
func (s *SellerService) RejectSeller(id uint, reason string, adminID uint) error {
	seller, err := s.sellerRepo.GetByID(id)
	if err != nil {
		return err
	}

	if seller.ApprovalStatus == "rejected" {
		return errors.New("seller is already rejected")
	}

	if reason == "" {
		return errors.New("rejection reason is required")
	}

	err = s.sellerRepo.ApproveReject(id, "rejected", reason, adminID)
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivityReject, "seller", &id, "Rejected seller registration: "+reason)

	return nil
}

// SuspendSeller suspends a seller account
func (s *SellerService) SuspendSeller(id uint, reason string, adminID uint) error {
	seller, err := s.sellerRepo.GetByID(id)
	if err != nil {
		return err
	}

	if seller.Status == constants.StatusSuspended {
		return errors.New("seller is already suspended")
	}

	err = s.sellerRepo.UpdateStatus(id, constants.StatusSuspended)
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivitySuspend, "seller", &id, "Suspended seller account: "+reason)

	return nil
}

// ReactivateSeller reactivates a suspended seller
func (s *SellerService) ReactivateSeller(id uint, adminID uint) error {
	seller, err := s.sellerRepo.GetByID(id)
	if err != nil {
		return err
	}

	if seller.Status != constants.StatusSuspended {
		return errors.New("seller is not suspended")
	}

	err = s.sellerRepo.UpdateStatus(id, "approved")
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivityActivate, "seller", &id, "Reactivated seller account")

	return nil
}

// SearchSellers searches sellers by query
func (s *SellerService) SearchSellers(query string, page, limit int) ([]*models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.sellerRepo.Search(query, page, limit)
}

// UpdateSeller updates seller information
func (s *SellerService) UpdateSeller(seller *models.User, adminID uint) error {
	err := s.sellerRepo.Update(seller)
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivityUpdate, "seller", &seller.ID, "Updated seller information")

	return nil
}

// GetSellerStats retrieves seller statistics
func (s *SellerService) GetSellerStats() (map[string]interface{}, error) {
	return s.sellerRepo.GetStats()
}

// GetTopSellers retrieves top sellers by sales
func (s *SellerService) GetTopSellers(limit int) ([]models.User, error) {
	if limit < 1 || limit > 50 {
		limit = 10
	}
	return s.sellerRepo.GetTopSellers(limit)
}

// Helper function to log activities
func (s *SellerService) logActivity(adminID uint, action, entityType string, entityID *uint, description string) {
	log := &models.AuditLog{
		AdminID:     &adminID,
		Action:      action,
		EntityType:  entityType,
		EntityID:    entityID,
		Description: description,
	}
	s.activityLogRepo.Create(log)
}
