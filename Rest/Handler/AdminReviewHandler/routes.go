package adminreviewhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *AdminReviewHandler) RegisterAdminReview(mux *http.ServeMux) {
	// Commission settings routes (admin only)

	mux.Handle("GET /api/admin/reviews",
		middleware.Chain(
			http.HandlerFunc(h.GetAllReviews),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("GET /api/admin/reviews/pending",
		middleware.Chain(
			http.HandlerFunc(h.GetPendingReviews),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)

	mux.Handle("GET /api/admin/reviews/reported",
		middleware.Chain(
			http.HandlerFunc(h.GetReportedReviews),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)

	mux.Handle("GET /api/admin/reviews/moderate",
		middleware.Chain(
			http.HandlerFunc(h.ModerateReview),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)

	mux.Handle("GET /api/admin/reviews/delete",
		middleware.Chain(
			http.HandlerFunc(h.DeleteReview),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)

	mux.Handle("GET /api/admin/stats",
		middleware.Chain(
			http.HandlerFunc(h.GetReviewStats),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)

	mux.Handle("POST /api/admin/reviews/report",
		middleware.Chain(
			http.HandlerFunc(h.ReportReview),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		),
	)
	mux.Handle("GET /api/admin/reviews/report",
		middleware.Chain(
			http.HandlerFunc(h.GetReportedReviews),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)
	mux.Handle("POST /api/admin/reviews/reports",
		middleware.Chain(
			http.HandlerFunc(h.GetReviewReports),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
			middleware.Requirerole([]string{"admin", "super_admin"}),
		),
	)

}
