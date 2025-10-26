package middleware

import (
	"fmt"
	"hifi/routes/rest"
	"net/http"
)

// -------------------- REWRITE --------------------

func RewriteRequest(r *http.Request) {

	r.URL.Scheme = "https"
	r.URL.Host = "api.tidal.com"
	r.Host = "tidal.com"

	q := r.URL.Query()

	print(q.Get("id"))
	switch r.URL.Path {
	case rest.Search3View():
		fmt.Println("recived")
	}
}
