package middleware

import (
	"hifi/config"
	"hifi/routes/rest"
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case rest.Search3View():
			w.Header().Set(config.HeaderAllowOrigin, "*")
		}
		next.ServeHTTP(w, r)
	})
}
