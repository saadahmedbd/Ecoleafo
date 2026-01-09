package reguserrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type RegUserRepository interface {
	FindByEmail(email string) (*models.RegUser, error)
	Create(user *models.RegUser) error
	Delete(id uint) error
}
type regUserRepository struct {
	db *gorm.DB
}

func NewRegUserRepository(db *gorm.DB) RegUserRepository {
	return &regUserRepository{
		db: db,
	}
}
