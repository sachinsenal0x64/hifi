package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/routes/rest"
	"hifi/types"
	"log/slog"
	"net/http"
	"net/url"
	"slices"
)

// -------------------- SESSION --------------------

func setQueryParams(q url.Values, params map[string]string) {
	for k, v := range params {
		q.Set(k, v)
	}
}

func Session(userName string, passWord string, ValidPaths []string) func(http.Handler) http.Handler {
	// proxy := httputil.NewSingleHostReverseProxy(target)

	// proxy.ModifyResponse = func(resp *http.Response) error {
	// 	resp.Header.Del("Access-Control-Allow-Origin")
	// 	return nil
	// }

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path == rest.Fresh() {
				next.ServeHTTP(w, r)
				return
			}

			if !slices.Contains(ValidPaths, r.URL.Path) {
				w.WriteHeader(config.StatusNotFound)
				return
			}

			if slices.Contains(ValidPaths, r.URL.Path) && r.URL.Path != rest.Ping() {
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

			// r.URL.Scheme = target.Scheme
			// r.URL.Host = target.Host
			// r.Host = target.Host

			ctx, store, err := Con()
			if err != nil {
				slog.Log(context.Background(), slog.LevelError, "Failed to connect to Valkey", "error", err)
			}
			defer store.Valkey.Close()

			// SET
			ok, err := store.Set(ctx, "cloud", "abc123")
			if err != nil {
				fmt.Println("SET error:", err)
			}
			fmt.Println("SET ok:", ok)

			// GET
			v, err := store.Get(ctx, "cloud")
			if err != nil {
				fmt.Println("GET error:", err)
			}
			fmt.Println("GET cloud:", v)

			// DEL
			// deleted, err := store.Del(ctx, "cloud")
			// if err != nil {
			// 	fmt.Println("DEL error:", err)
			// }
			// fmt.Println("DEL cloud:", deleted)

			v3, err := store.Get(ctx, "cloud")
			if err != nil {
				fmt.Println("GET error:", err)
			}
			fmt.Println("GET cloud:", v3)

			// Mock Subsonic response for /ping endpoint

			var resp types.SubsonicWrapper
			resp.Subsonic.Status = "ok"
			resp.Subsonic.Version = "1.15.0"
			resp.Subsonic.Type = "hifi"
			resp.Subsonic.ServerVersion = "0.19.0"
			resp.Subsonic.OpenSubsonic = true

			w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
			w.WriteHeader(http.StatusOK)

			b, err := json.Marshal(resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = w.Write(b)

			/* Forward the request to the
			subsonic server -> (gonic) */

			// proxy.ServeHTTP(w, r)
		})
	}
}
