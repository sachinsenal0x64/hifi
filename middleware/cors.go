package middleware

import (
	"hifi/config"
	"net/http"
	"slices"
	"strings"
)

// -------------------- CORS --------------------

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Ensure path ends with .view (for Tempo)
		if !strings.HasSuffix(path, ".view") {
			path += ".view"
		}

		if slices.Contains(config.ValidPaths, path) {
			w.Header().Set(config.HeaderAllowOrigin, config.CORSAllowOrigin)
		}
		next.ServeHTTP(w, r)
	})
}
