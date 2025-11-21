package middleware

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/constants"
)

func Requirerole(allowedRoles []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roleVal := r.Context().Value(constants.ContextKeyRole)
			if roleVal == nil {
				http.Error(w, "Unauthorized: no role in context", http.StatusUnauthorized)
				return
			}

			var role string
			switch v := roleVal.(type) {
			case string:
				role = v
			case []interface{}:
				if len(v) > 0 {
					if s, ok := v[0].(string); ok {
						role = s
					}
				}
			default:
				http.Error(w, "Unauthorized: invalid role type", http.StatusUnauthorized)
				return
			}

			// Check if role is allowed
			authorized := false
			for _, ar := range allowedRoles {
				if role == ar {
					authorized = true
					break
				}
			}
			if !authorized {
				http.Error(w, "Unauthorized: role not allowed", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
