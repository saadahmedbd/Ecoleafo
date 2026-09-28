package middleware

import (
	"net/http"
	"os"
	"strings"
)

func isAllowedOrigin(origin string) bool {
	if origin == "" {
		return false
	}

	// Always allow local development
	if strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:") {
		return true
	}

	// Allow production ecoleafo domains
	if origin == "https://ecoleafo.com" || origin == "https://www.ecoleafo.com" || origin == "http://ecoleafo.com" || origin == "http://www.ecoleafo.com" {
		return true
	}

	// Allow any Vercel deployment/preview (*.vercel.app)
	if strings.HasSuffix(origin, ".vercel.app") {
		return true
	}

	// Allow origins defined in ALLOWED_ORIGINS env variable
	envOrigins := os.Getenv("ALLOWED_ORIGINS")
	if envOrigins != "" {
		for _, o := range strings.Split(envOrigins, ",") {
			if strings.TrimSpace(o) == origin {
				return true
			}
		}
	}

	return false
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if isAllowedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Expose-Headers", "X-Token-Expired, Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
