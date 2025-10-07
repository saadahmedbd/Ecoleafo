package selleraccountrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type SellerRegistrationRepository interface {
	// registration
	CreateRegUser(regUser *models.RegUser) error
	CreateSeller(seller *models.User) error
	GetSellerByEmail(email string) (*models.User, error)
	GetSellerByUserId(userID uint) (*models.User, error)

	//profile management
	UpdateSellerProfile(seller *models.User) error
	GetSellerWithRelation(sellerID uint) (*models.User, error)
	CheckProfileCompletion(sellerID uint) (bool, []string, error)

	//payment method
	CreatePaymentMethod(paymentMethod *models.SellerPaymentMethod) error
	GetPaymentMethods(sellerID uint) ([]models.SellerPaymentMethod, error)
	GetPaymentMethodByID(paymentMethodID, sellerID uint) (*models.SellerPaymentMethod, error)
	UpdatePaymentMethod(paymentMethod *models.SellerPaymentMethod) error
	DeletePaymentMethod(paymentMethodID, sellerID uint) error
	SetDefaultPaymentMethod(sellerID, paymentMethodID uint) error
}
type sellerRegistrationRepository struct {
	db *gorm.DB
}

func NewSellerRegistrationRepository(db *gorm.DB) SellerRegistrationRepository {
	return &sellerRegistrationRepository{db: db}
}
