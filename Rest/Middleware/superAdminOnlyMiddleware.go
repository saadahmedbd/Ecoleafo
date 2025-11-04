package middleware

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// SuperAdminOnlyMiddleware ensures only super admins can access the route
func SuperAdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(constants.ContextKeyRole).(string)

		if role != constants.RoleSuperAdmin {
			util.SendError(w, "Super admin access required", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
