package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, `{"error":"Missing Authorization header"}`, http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, `{"error":"Missing or invalid auth header"}`, http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
		if tokenStr == "" {
			http.Error(w, `{"error":"Missing token after Bearer"}`, http.StatusUnauthorized)
			return
		}

		claims, err := util.VerifyJwt(tokenStr)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid or expired token: %s"}`, err.Error()), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		
		// Add user_id and role to headers for easy access
		if userID, ok := claims["user_id"]; ok {
			r.Header.Set("user_id", fmt.Sprintf("%v", userID))
		}
		if roles, ok := claims["role"]; ok {
			r.Header.Set("user_role", fmt.Sprintf("%v", roles))
		}
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// func AuthenticateJWT(next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		header := r.Header.Get("Authorization")
// 		if header == "" {
// 			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
// 			return
// 		}
// 		if !strings.HasPrefix(header, "Bearer ") {
// 			http.Error(w, "missing or invalid auth header", http.StatusUnauthorized)
// 			return
// 		}
// 		tokenStr := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
// 		if tokenStr == "" {
// 			http.Error(w, "Missing token after Bearer", http.StatusUnauthorized)
// 			return
// 		}

// 		claims, err := util.VerifyJwt(tokenStr)
// 		if err != nil {
// 			http.Error(w, "Invalid or expired token: "+err.Error(), http.StatusUnauthorized)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "claims", claims)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// type key string

// const UserIDKey key = "user_id"

// func AuthenticateJWT(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		header := r.Header.Get("Authorization")
// 		if header == "" {
// 			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
// 			return
// 		}

// 		tokenStr := strings.TrimPrefix(header, "Bearer ")

// 		cnf := Config.GetConfig()
// 		payload, err := util.VerifyJwt(tokenStr, cnf.JwtSecretKey) // custom JWT verification
// 		if err != nil {
// 			http.Error(w, "Invalid or expired token: "+err.Error(), http.StatusUnauthorized)
// 			return
// 		}

// 		// save claims into context for the handler to use
// 		ctx := context.WithValue(r.Context(), "claims", map[string]interface{}{
// 			"user_id":    payload.UserId,
// 			"role":       payload.Role,
// 			"first_name": payload.FirstName,
// 			"last_name":  payload.LastName,
// 		})
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
