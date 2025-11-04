package adminmangementhandler

import (
	"fmt"
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *DashboardHandler) RegisterDashboardRoutes(mux *http.ServeMux) {
	fmt.Println("Registering dashboard route: GET /api/dashboard/stats")
	
	// Test route without middleware
	mux.HandleFunc("GET /api/dashboard/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "message": "Dashboard route is working"}`))
	})
	
	mux.Handle("GET /api/dashboard/stats", middleware.Chain(http.HandlerFunc(h.GetDashboardStats),
		middleware.Logger,
		middleware.Cors,
		middleware.AuthenticateJWT,
		middleware.AdminOnlyMiddleware,
	))
}
