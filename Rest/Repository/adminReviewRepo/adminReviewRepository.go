package adminreviewrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
	"gorm.io/gorm"
)

type AdminReviewRepository interface {
	//admin operation
	GetAllReviews(query adminreviewdto.ReviewListQuery) ([]models.Review, int64, error)
	GetReviewByID(id uint) (*models.Review, error)
	GetPendingReviews(page, limit int) ([]models.Review, int64, error)
	GetReportedReviews(page, limit int) ([]models.Review, int64, error)
	ApproveReview(id uint, adminID uint) error
	RejectReview(id uint, reason string, adminID uint) error
	DeleteReview(id uint) error
	// Reports
	GetReviewReports(status string, page, limit int) ([]models.ReviewReport, int64, error)
	GetReportsByReviewID(reviewID uint) ([]models.ReviewReport, error)
	CreateReviewReport(report *models.ReviewReport) error
	UpdateReportStatus(reportID uint, status string, adminID uint) error

	// Statistics
	GetReviewStats() (map[string]interface{}, error)
	GetProductReviewStats(productID uint) (map[string]interface{}, error)
}
type adminReviewRepository struct {
	db *gorm.DB
}

func NewAdminReviewRepository(db *gorm.DB) AdminReviewRepository {
	return &adminReviewRepository{db: db}
}
