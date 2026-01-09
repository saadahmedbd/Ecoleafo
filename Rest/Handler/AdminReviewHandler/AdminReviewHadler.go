package adminreviewhandler

import adminreviewservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminReviewService"

type AdminReviewHandler struct {
	adminreviewservice adminreviewservice.AdminReviewService
}

func NewAdminReviewService(adminreviewservice adminreviewservice.AdminReviewService) *AdminReviewHandler {
	return &AdminReviewHandler{
		adminreviewservice: adminreviewservice,
	}
}
