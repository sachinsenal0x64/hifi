package middleware

import (
	"hifi/config"
	"hifi/routes/rest"
	"log/slog"
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

			// Blacklist URL paths
			if slices.Contains(exclude, r.URL.Path) {
				w.WriteHeader(config.StatusNotFound)
				return
			}

			// Rewrite tidal requests
			switch r.URL.Path {
			case rest.Search3View(), rest.GetArtistsView(),
				rest.GetCoverArtView(), rest.Stream(),
				rest.GetSong(), rest.Scrobble(),
				rest.GetAlbumView():
				RewriteRequest(w, r)
				return

			}

			/* Add authentication parameters
			to the URL query like -> (https://) */

			q := r.URL.Query()

			s := q.Get("s")
			t := q.Get("t")
			f := q.Get("f")
			c := q.Get("c")

			userName := q.Get("u")
			passWord := q.Get("p")

			// -------------------- SESSION --------------------

			// salt := Salt("Key") //random string
			// token := Token("Password", salt)

			params := map[string]string{
				"u": userName,
				"c": c,
				"f": f,
			}

			// Check if s and t exist in query

			if s != "" && t != "" {
				// Use token authentication
				params["s"] = s
				params["t"] = t
			} else {
				/* Fallback to legacy password
				authentication */
				params["p"] = passWord
			}

			setQueryParams(q, params)

			r.URL.RawQuery = q.Encode()
			slog.Info("incoming request",
				"path", r.URL.Path,
				"raw", r.URL.RawQuery,
			)

			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.Host = target.Host

			/* Forward the request to the
			subsonic server -> (gonic) */
			proxy.ServeHTTP(w, r)
		})
	}
}
