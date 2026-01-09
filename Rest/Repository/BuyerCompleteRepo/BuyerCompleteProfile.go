package buyercompleterepo

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"

	"gorm.io/gorm"
)

type ProfileCompleteRepository interface {
	GetBuyerByUserID(userID uint) (*models.Buyer, error)
	UpdateBuyerProfile(buyer *models.Buyer) error
	CreateAddress(address *models.Address) error
	GetBuyerAddresses(buyerID uint) ([]models.Address, error)
	HasCompletedProfile(userID uint) (bool, []string, error)
}

type profileCompleteRepository struct {
	db *gorm.DB
}

func NewProfileCompleteRepository(db *gorm.DB) ProfileCompleteRepository {
	return &profileCompleteRepository{
		db: db,
	}
}
func (r *profileCompleteRepository) GetBuyerByUserID(userID uint) (*models.Buyer, error) {
	var buyer models.Buyer
	err := r.db.Where("user_id = ?", userID).
		First(&buyer).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("buyer not found")
		}
		return nil, err
	}
	return &buyer, nil
}
func (r *profileCompleteRepository) UpdateBuyerProfile(buyer *models.Buyer) error {
	return r.db.Save(buyer).Error
}
func (r *profileCompleteRepository) CreateAddress(address *models.Address) error {
	//if this is set as default, unset other
	if address.IsDefault {
		r.db.Model(&models.Address{}).
			Where("buyer_id = ?", address.BuyerID).
			Update("is_default", false)
	}
	return r.db.Create(address).Error
}

func (r *profileCompleteRepository) GetBuyerAddresses(buyerID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.Where("buyer_id = ?", buyerID).
		Order("is_default DESC, created_at DESC").
		Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}
func (r *profileCompleteRepository) HasCompletedProfile(userID uint) (bool, []string, error) {
	var buyer models.Buyer
	err := r.db.Where("user_id = ?", userID).First(&buyer).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// If buyer doesn't exist, profile is incomplete
			return false, []string{"Phone", "Address"}, nil
		}
		return false, nil, err
	}
	var MissingFields []string
	//check phone
	if buyer.Phone == "" {
		MissingFields = append(MissingFields, "Phone")
	}
	//check if has at least one address
	var addressCount int64
	r.db.Model(&models.Address{}).Where("buyer_id = ?", buyer.ID).Count(&addressCount)
	if addressCount == 0 {
		MissingFields = append(MissingFields, "Address")
	}
	return len(MissingFields) == 0, MissingFields, nil
}
