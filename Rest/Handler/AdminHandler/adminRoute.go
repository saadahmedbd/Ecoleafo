package adminhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Adminhandler) AdminRoute(mux *http.ServeMux) {
	//public routes
	mux.Handle("POST /api/admin/register", middleware.Chain(http.HandlerFunc(h.RegisterAdmin),
		middleware.Logger,
	))
	mux.Handle("GET /api/admin/validate-invitation", middleware.Chain(http.HandlerFunc(h.ValidateInvitation),
		middleware.Logger,
	))

	//protected route auth recuired
	mux.Handle("GET /api/admin/profile", middleware.Chain(http.HandlerFunc(h.GetAdminProfile),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/admin/get", middleware.Chain(http.HandlerFunc(h.GetAdminByID),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/admin/all", middleware.Chain(http.HandlerFunc(h.GetAllAdmins),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("PUT /api/admin/update", middleware.Chain(http.HandlerFunc(h.UpdateAdmin),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))

	// super admin route
	mux.Handle("POST /api/admin/invite", middleware.Chain(http.HandlerFunc(h.CreateInvitation),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/admin/update-permissions", middleware.Chain(http.HandlerFunc(h.UpdatePermission),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))

	mux.Handle("POST /api/admin/deactivate", middleware.Chain(http.HandlerFunc(h.DeactivateAdmin),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/admin/activate", middleware.Chain(http.HandlerFunc(h.ActivateAdmin),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/admin/invitations", middleware.Chain(http.HandlerFunc(h.GetPendingInvitations),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("DELETE /api/admin/cancel-invitation", middleware.Chain(http.HandlerFunc(h.CancelInvitation),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))

}
