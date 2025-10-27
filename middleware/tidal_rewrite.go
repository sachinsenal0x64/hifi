package middleware

import (
	"hifi/config"
	"hifi/routes/rest"
	"net/http"
)

// -------------------- REWRITE --------------------

func RewriteRequest(r *http.Request) {

	r.URL.Scheme = config.Scheme
	r.URL.Host = config.TidalHost
	r.Host = config.BrowserHost

	q := r.URL.Query()

	switch r.URL.Path {
	case rest.Search3View():
		print("query:", q.Get("query"))
	}
}
