package middleware

import (
	"hifi/config"
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

			// Add authentication parameters to the URL query (https://)
			q := r.URL.Query()
			setQueryParams(q, map[string]string{
				"message": "hi",
				"u":       userName,
				"p":       passWord,
				"c":       "",
				"f":       "json",
			})

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
