package middleware

import (
	"hifi/config"
	"net/http"
	"slices"
)

// -------------------- CORS --------------------

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(config.ValidPaths, r.URL.Path) {
			w.Header().Set(config.HeaderAllowOrigin, config.CORSAllowOrigin)
		}
		next.ServeHTTP(w, r)
	})
}
