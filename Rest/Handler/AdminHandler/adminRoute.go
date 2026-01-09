package adminhandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Adminhandler) AdminRoute(mux *http.ServeMux) {
	//public routes
	mux.Handle("POST /api/admin/register", middleware.Chain(http.HandlerFunc(h.RegisterAdmin),
		middleware.Logger,
		middleware.Cors,
	))
	mux.Handle("GET /api/admin/validate-invitation", middleware.Chain(http.HandlerFunc(h.ValidateInvitation),
		middleware.Logger,
	))

	//protected route auth recuired
	mux.Handle("GET /api/admin/profile", middleware.Chain(http.HandlerFunc(h.GetAdminProfile),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("GET /api/admin/get", middleware.Chain(http.HandlerFunc(h.GetAdminByID),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("GET /api/admin/all", middleware.Chain(http.HandlerFunc(h.GetAllAdmins),

		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("PUT /api/admin/update", middleware.Chain(http.HandlerFunc(h.UpdateAdmin),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))

	// super admin route
	mux.Handle("POST /api/admin/invite", middleware.Chain(http.HandlerFunc(h.CreateInvitation),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("POST /api/admin/update-permissions", middleware.Chain(http.HandlerFunc(h.UpdatePermission),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))

	mux.Handle("POST /api/admin/deactivate", middleware.Chain(http.HandlerFunc(h.DeactivateAdmin),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("POST /api/admin/activate", middleware.Chain(http.HandlerFunc(h.ActivateAdmin),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("GET /api/admin/invitations", middleware.Chain(http.HandlerFunc(h.GetPendingInvitations),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))
	mux.Handle("DELETE /api/admin/cancel-invitation", middleware.Chain(http.HandlerFunc(h.CancelInvitation),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.Requirerole([]string{"admin", "super_admin"}),
	))

}
