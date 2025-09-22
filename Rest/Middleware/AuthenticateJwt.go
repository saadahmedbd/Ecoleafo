package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/saadahmedbd/Treestore/Config"
	util "github.com/saadahmedbd/Treestore/Util"
)

type key string

const UserIDKey key = "user_id"

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")

		cnf := Config.GetConfig()
		payload, err := util.VerifyJwt(tokenStr, cnf.JwtSecretKey) // custom JWT verification
		if err != nil {
			http.Error(w, "Invalid or expired token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// save claims into context for the handler to use
		ctx := context.WithValue(r.Context(), "claims", map[string]interface{}{
			"user_id":    payload.UserId,
			"role":       payload.Role,
			"first_name": payload.FirstName,
			"last_name":  payload.LastName,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
