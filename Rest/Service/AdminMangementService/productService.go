package adminmangementservice

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	adminmangement "github.com/saadahmedbd/Treestore/Rest/Repository/AdminMangement"
	"github.com/saadahmedbd/Treestore/constants"
)

type ProductService struct {
	productRepo     *adminmangement.ProductRepository
	activityLogRepo *adminmangement.AuditLogRepository
}

func NewProductService(productRepo *adminmangement.ProductRepository, activityLogRepo *adminmangement.AuditLogRepository) *ProductService {
	return &ProductService{
		productRepo:     productRepo,
		activityLogRepo: activityLogRepo,
	}
}

// GetAllProducts retrieves all products with pagination
func (s *ProductService) GetAllProducts(page, limit int, status string) ([]models.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.productRepo.GetAll(page, limit, status)
}

// GetProductByID retrieves a product by ID
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepo.GetByID(id)
}

// GetPendingApprovals retrieves products pending approval
func (s *ProductService) GetPendingApprovals(page, limit int) ([]models.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.productRepo.GetPendingApprovals(page, limit)
}

// ApproveProduct approves a product listing
func (s *ProductService) ApproveProduct(id uint, adminID uint) error {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	if product.ApprovalStatus == "approved" {
		return errors.New("product is already approved")
	}

	err = s.productRepo.ApproveProduct(id, adminID)
	if err != nil {
		return err
	}

	// Log activity
	s.logProductActivity(adminID, constants.ActionApprove, &id, "Approved product listing")

	return nil
}

// RejectProduct rejects a product listing
func (s *ProductService) RejectProduct(id uint, reason string, adminID uint) error {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	if product.ApprovalStatus == "rejected" {
		return errors.New("product is already rejected")
	}

	if reason == "" {
		return errors.New("rejection reason is required")
	}

	err = s.productRepo.ApproveReject(id, "rejected", reason, adminID)
	if err != nil {
		return err
	}

	// Log activity
	s.logProductActivity(adminID, constants.ActionReject, &id, "Rejected product listing: "+reason)

	return nil
}

// SearchProducts searches products by query
func (s *ProductService) SearchProducts(query string, page, limit int) ([]models.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.productRepo.Search(query, page, limit)
}

// UpdateProduct updates product information
func (s *ProductService) UpdateProduct(product *models.Product, adminID uint) error {
	err := s.productRepo.Update(product)
	if err != nil {
		return err
	}

	// Log activity
	s.logProductActivity(adminID, constants.ActionUpdate, &product.ID, "Updated product information")

	return nil
}

// DeleteProduct soft deletes a product
func (s *ProductService) DeleteProduct(id uint, adminID uint) error {
	err := s.productRepo.Delete(id)
	if err != nil {
		return err
	}

	// Log activity
	s.logProductActivity(adminID, constants.ActionDelete, &id, "Deleted product")

	return nil
}

// GetProductStats retrieves product statistics
func (s *ProductService) GetProductStats() (map[string]interface{}, error) {
	return s.productRepo.GetStats()
}

// Helper function
func (s *ProductService) logProductActivity(adminID uint, action string, entityID *uint, description string) {
	log := &models.AuditLog{
		ActorID:     &adminID,
		Action:      action,
		EntityType:  "product",
		EntityID:    entityID,
		Description: description,
	}
	s.activityLogRepo.Create(log)
}
