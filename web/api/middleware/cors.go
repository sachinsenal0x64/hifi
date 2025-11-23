package middleware

import (
	"api/config"
	"fmt"
	"net/http"
	"slices"
)

// -------------------- CORS --------------------

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(config.ValidPaths, r.URL.Path) {
			w.Header().Set(config.HeaderAllowOrigin, config.CORSAllowOrigin)
			w.Header().Set(config.HeaderContentSecurityPolicy, fmt.Sprintf("script-src %s", config.Scheme+"://"+config.ProxyHost))
		}

		next.ServeHTTP(w, r)
	})
}
