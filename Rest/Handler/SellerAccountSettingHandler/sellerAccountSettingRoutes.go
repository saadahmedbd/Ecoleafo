package selleraccountsettinghandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Selleraccountsettinghandler) RegisterSellerAccountSettingHandler(mux *http.ServeMux) {
	//profile
	mux.Handle("GET /api/seller/profile", middleware.Chain(http.HandlerFunc(h.GetFullProfile),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/seller/statistics", middleware.Chain(http.HandlerFunc(h.GetStatistics),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//account managment
	mux.Handle("PUT /api/seller/account", middleware.Chain(http.HandlerFunc(h.UpdateAccount),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/seller/account/profile", middleware.Chain(http.HandlerFunc(h.UploadProfilePhoto),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/seller/password/change", middleware.Chain(http.HandlerFunc(h.ChangePassword),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/seller/account/deactivate", middleware.Chain(http.HandlerFunc(h.DeactivateAccount),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("DELETE /api/seller/account/delete", middleware.Chain(http.HandlerFunc(h.DeleteAccount),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//store management
	mux.Handle("PUT /api/seller/store", middleware.Chain(http.HandlerFunc(h.UpdateStore),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("PUT /api/seller/store/branding", middleware.Chain(http.HandlerFunc(h.UpdateBranding),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("PUT /api/seller/store/policies", middleware.Chain(http.HandlerFunc(h.UpdatePolicies),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//notification
	mux.Handle("GET /api/seller/notifications/preferences", middleware.Chain(http.HandlerFunc(h.GetNotificationPreferences),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("PUT /api/seller/notifications/preferences", middleware.Chain(http.HandlerFunc(h.UpdateNotificationPreferences),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//verification
	mux.Handle("GET /api/seller/verification/status", middleware.Chain(http.HandlerFunc(h.GetVerificationStatus),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /api/seller/verification/upload", middleware.Chain(http.HandlerFunc(h.UploadVerificationDocument),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	//security
	mux.Handle("PUT /api/seller/security/2fa", middleware.Chain(http.HandlerFunc(h.Toggle2FA),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /api/seller/security/activity", middleware.Chain(http.HandlerFunc(h.GetLoginActivity),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
	))

}
