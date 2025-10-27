package middleware

import (
	"hifi/config"
	"hifi/routes/rest"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"slices"
)

// -------------------- SESSION --------------------

func setQueryParams(q url.Values, params map[string]string) {
	for k, v := range params {
		q.Set(k, v)
	}
}

func Session(userName, passWord, targetHost string, exclude []string) func(http.Handler) http.Handler {
	target, _ := url.Parse(targetHost)
	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if slices.Contains(exclude, r.URL.Path) {
				w.WriteHeader(config.StatusNotFound)
				return
			}

			switch r.URL.Path {
			case rest.Search3View():
				RewriteRequest(r)
			}

			// Add authentication parameters to the URL query like -> (https://)
			q := r.URL.Query()

			// q.Del("s")
			// q.Del("t")

			params := map[string]string{
				"u": userName,
				"c": "",
				"f": "json",
			}

			// Check if s and t exist in query
			s := q.Get("s")
			t := q.Get("t")

			if s != "" && t != "" {

				salt := Encode(16)

				// Token auth exists
				params["s"] = salt
				params["t"] = "fd0ca358286130b31587f1cd2f0ddf7c"
			} else {
				// Fallback to legacy password
				params["p"] = passWord
			}

			setQueryParams(q, params)

			r.URL.RawQuery = q.Encode()
			log.Println("PATH:", r.URL.Path, "RAW:", r.URL.RawQuery)

			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.Host = target.Host

			// Forward the request to the subsonic server (gonic)
			proxy.ServeHTTP(w, r)
		})
	}
}
