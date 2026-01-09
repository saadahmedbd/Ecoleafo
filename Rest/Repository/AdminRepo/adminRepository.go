package adminrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type AdminRepo interface {
	Create(admin *models.Admin) error
	FindByID(id uint) (*models.Admin, error)
	FindByUserID(userID uint) (*models.Admin, error)
	FindAll(page, limit int) ([]models.Admin, int64, error)
	Update(admin *models.Admin) error
	Delete(id uint) error
	UpdateLastLogin(userID uint) error

	// Invitation methods
	CreateInvitation(invitation *models.AdminInvitation) error
	FindInvitationByToken(token string) (*models.AdminInvitation, error)
	FindInvitationByEmail(email string) (*models.AdminInvitation, error)
	GetPendingInvitations() ([]models.AdminInvitation, error)
	UpdateInvitation(invitation *models.AdminInvitation) error
	DeleteInvitation(id uint) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepo {
	return &adminRepository{
		db: db,
	}
}
