package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, `{"error":"Missing or invalid auth header"}`, http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
		claims, err := util.VerifyJwt(tokenStr)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid or expired token: %s"}`, err.Error()), http.StatusUnauthorized)
			return
		}

		// Extract user_id
		userID, ok1 := claims["user_id"].(float64)
		if !ok1 {
			util.SendError(w, "Invalid token claims structure", http.StatusUnauthorized)
			return
		}

		// Extract role (support both single string or array)
		roleVal, ok2 := claims["role"]
		if !ok2 || roleVal == nil {
			util.SendError(w, "Invalid token claims structure", http.StatusUnauthorized)
			return
		}

		var roleStr string
		switch v := roleVal.(type) {
		case string:
			roleStr = v
		case []interface{}:
			if len(v) > 0 {
				if s, ok := v[0].(string); ok {
					roleStr = s
				}
			}
		}

		//  Add to context
		ctx := context.WithValue(r.Context(), constants.ContextKeyUserID, uint(userID))
		ctx = context.WithValue(ctx, constants.ContextKeyRole, roleVal)

		// Also set headers for backward compatibility
		r.Header.Set("user_id", fmt.Sprintf("%.0f", userID))
		r.Header.Set("user_role", roleStr)

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
