package middleware

import (
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/routes/rest"
	"hifi/types"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"slices"

	"github.com/alexedwards/argon2id"
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

			// salt := Salt(config.Salt)
			// token := Token(passWord, salt)
			// _ = token

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

			// Mock Subsonic response for /ping endpoint

			url := fmt.Sprintf("%s://%s/v1/secrets/?env=%s&path=/hifi_users&key=%s&app_id=%s", config.Http, config.ProxyHost, config.ENV, userName, config.AppID)

			slog.Info(url)

			method := "GET"

			client := &http.Client{}
			req, err := http.NewRequest(method, url, nil)

			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Add("Authorization", "Bearer User "+config.ProxyKey)
			req.Header.Add(config.HeaderContentType, config.ContentTypeJSON)

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			var g types.AppGet

			if err := json.Unmarshal(body, &g); err != nil {
				slog.Error("error unmarshalling response", "error", err)
				return
			}

			match, err := argon2id.ComparePasswordAndHash(passWord, g[0].Password)
			if err != nil {
				writeSubsonic(w, "failed", http.StatusBadRequest)
			}

			if !match {
				writeSubsonic(w, "failed", http.StatusBadRequest)
			} else {
				writeSubsonic(w, "ok", http.StatusOK)
			}

			/* Forward the request to the
			subsonic server -> */

			// proxy.ServeHTTP(w, r)
		})
	}
}
