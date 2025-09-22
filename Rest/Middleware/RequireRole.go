package middleware

import (
	"net/http"
	
)

func Requirerole(allowedRoles []string, next http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get role from context (set by AuthenticateJWT middleware)
		//claims ==payload
		claims, ok := r.Context().Value("claims").(map[string]interface{})
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		role, ok := claims["role"].(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Check if user has required role
		authorized := false
		for _, ar := range allowedRoles {
			if role == ar {
				authorized = true
				break
			}
		}
		if !authorized {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}
