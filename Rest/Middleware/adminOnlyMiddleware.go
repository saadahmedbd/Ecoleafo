package middleware

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// AdminOnlyMiddleware ensures only admins can access the route
func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Safely get the role from context
		roleVal := r.Context().Value(constants.ContextKeyRole)
		if roleVal == nil {
			util.SendError(w, "Missing role in context", http.StatusUnauthorized)
			return
		}

		// Handle both string and []interface{} roles
		switch roles := roleVal.(type) {
		case string:
			if roles != constants.RoleAdmin && roles != constants.RoleSuperAdmin {
				util.SendError(w, "Admin access required", http.StatusForbidden)
				return
			}

		case []interface{}:
			isAdmin := false
			for _, r := range roles {
				if str, ok := r.(string); ok && (str == constants.RoleAdmin || str == constants.RoleSuperAdmin) {
					isAdmin = true
					break
				}
			}
			if !isAdmin {
				util.SendError(w, "Admin access required", http.StatusForbidden)
				return
			}

		default:
			util.SendError(w, "Invalid role format", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
