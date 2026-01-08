package reviewrepo

import "gorm.io/gorm"

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

// GetDB returns the database instance
func (r *ReviewRepository) GetDB() *gorm.DB {
	return r.db
}
