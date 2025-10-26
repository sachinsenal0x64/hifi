package middleware

import (
	"fmt"
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
	case "/rest/search3.view":
		fmt.Println("recived")
	}
}
