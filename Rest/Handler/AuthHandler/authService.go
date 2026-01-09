package authhandler

import "gorm.io/gorm"

type authService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *authService {
	return &authService{db: db}
}
