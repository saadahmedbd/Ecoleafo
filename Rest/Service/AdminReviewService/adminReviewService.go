package adminreviewservice

import (
	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
	auditlogrepo "github.com/saadahmedbd/Treestore/Rest/Repository/AuditLogRepo"
	productrepo "github.com/saadahmedbd/Treestore/Rest/Repository/ProductRepo"
	adminreviewrepo "github.com/saadahmedbd/Treestore/Rest/Repository/adminReviewRepo"
)

type AdminReviewService interface {
	// Admin Operations
	GetAllReviews(query adminreviewdto.ReviewListQuery) ([]adminreviewdto.ReviewResponse, int64, error)
	GetPendingReviews(page, limit int) ([]adminreviewdto.ReviewResponse, int64, error)
	GetReportedReviews(page, limit int) ([]adminreviewdto.ReviewResponse, int64, error)
	ModerateReview(id uint, req *adminreviewdto.ModerateReviewRequest, adminID uint) error
	DeleteReview(id uint, adminID uint) error

	// Reports
	GetReviewReports(status string, page, limit int) ([]adminreviewdto.ReviewReportResponse, int64, error)
	ReportReview(reviewID uint, req *adminreviewdto.ReportReviewRequest, reporterID uint, reporterType string) error
	ReviewReport(reportID uint, action string, adminID uint) error

	// Statistics
	GetReviewStats() (*adminreviewdto.ReviewStatsResponse, error)
}
type adminReviewService struct {
	adminreviewRepo adminreviewrepo.AdminReviewRepository
	productRepo     productrepo.ProductRepository
	auditlogRepo    auditlogrepo.AuditLogRepository
}

func NewAdminReviewService(adminreviewrepo adminreviewrepo.AdminReviewRepository, productrepo productrepo.ProductRepository, auditlogrepo auditlogrepo.AuditLogRepository) AdminReviewService {
	return &adminReviewService{
		adminreviewRepo: adminreviewrepo,
		productRepo:     productrepo,
		auditlogRepo:    auditlogrepo,
	}
}
