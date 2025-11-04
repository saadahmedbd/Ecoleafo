package adminmangementservice

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	adminmangement "github.com/saadahmedbd/Treestore/Rest/Repository/AdminMangement"
	"github.com/saadahmedbd/Treestore/constants"
)

type BuyerService struct {
	buyerRepo    *adminmangement.BuyerRepository
	auditLogRepo *adminmangement.AuditLogRepository
}

func NewBuyerService(buyerRepo *adminmangement.BuyerRepository, auditLogRepo *adminmangement.AuditLogRepository) *BuyerService {
	return &BuyerService{
		buyerRepo:    buyerRepo,
		auditLogRepo: auditLogRepo,
	}

}

// GetAllBuyers retrieves all Buyers with pagination
func (s *BuyerService) GetAllBuyers(page, limit int, status string) ([]models.Buyer, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.buyerRepo.GetAll(page, limit, status)
}

// GetBuyerByID retrieves a Buyer by ID
func (s *BuyerService) GetBuyerByID(id uint) (*models.Buyer, error) {
	return s.buyerRepo.GetByID(id)
}

// SearchBuyers searches Buyers by query
func (s *BuyerService) SearchBuyers(query string, page, limit int) ([]models.Buyer, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.buyerRepo.Search(query, page, limit)
}

// ActivateBuyer activates a Buyer account
func (s *BuyerService) ActivateBuyer(id uint, adminID uint) error {
	Buyer, err := s.buyerRepo.GetByID(id)
	if err != nil {
		return err
	}

	if Buyer.Status == constants.StatusActive {
		return errors.New("Buyer is already active")
	}

	err = s.buyerRepo.UpdateStatus(id, constants.StatusActive)
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivityActivate, "Buyer", &id, "Activated Buyer account")

	return nil
}

// DeactivateBuyer deactivates a Buyer account
func (s *BuyerService) DeactivateBuyer(id uint, adminID uint) error {
	Buyer, err := s.buyerRepo.GetByID(id)
	if err != nil {
		return err
	}

	if Buyer.Status == constants.StatusInactive {
		return errors.New("Buyer is already inactive")
	}

	err = s.buyerRepo.UpdateStatus(id, constants.StatusInactive)
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivityDeactivate, "Buyer", &id, "Deactivated Buyer account")

	return nil
}

// SuspendBuyer suspends a Buyer account
func (s *BuyerService) SuspendBuyer(id uint, adminID uint) error {
	Buyer, err := s.buyerRepo.GetByID(id)
	if err != nil {
		return err
	}

	if Buyer.Status == constants.StatusSuspended {
		return errors.New("Buyer is already suspended")
	}

	err = s.buyerRepo.UpdateStatus(id, constants.StatusSuspended)
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivitySuspend, "Buyer", &id, "Suspended Buyer account")

	return nil
}

// UpdateBuyer updates Buyer information
func (s *BuyerService) UpdateBuyer(Buyer *models.Buyer, adminID uint) error {
	err := s.buyerRepo.Update(Buyer)
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivityUpdate, "Buyer", &Buyer.ID, "Updated Buyer information")

	return nil
}

// DeleteBuyer soft deletes a Buyer
func (s *BuyerService) DeleteBuyer(id uint, adminID uint) error {
	err := s.buyerRepo.Delete(id)
	if err != nil {
		return err
	}

	// Log activity
	s.logActivity(adminID, constants.ActivityDelete, "Buyer", &id, "Deleted Buyer account")

	return nil
}

// GetBuyerStats retrieves Buyer statistics
func (s *BuyerService) GetBuyerStats() (map[string]int64, error) {
	return s.buyerRepo.GetStats()
}

// Helper function to log activities
func (s *BuyerService) logActivity(adminID uint, action, entityType string, entityID *uint, description string) {
	log := &models.AuditLog{
		AdminID:     &adminID,
		Action:      action,
		EntityType:  entityType,
		EntityID:    entityID,
		Description: description,
	}
	s.auditLogRepo.Create(log)
}
