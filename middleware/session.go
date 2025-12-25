// Do not delete commented code

package middleware

import (
	"encoding/hex"
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
	"strings"

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

	// Define internal mock paths that don't need the Rewrite logic but need valid JSON responses
	mockPaths := []string{
		rest.GetMusicFoldersView(),
		"/rest/getLicense.view",
		"/rest/getPlaylists.view",
		"/rest/getGenres.view",
		"/rest/getStarred.view",
		"/rest/getStarred2.view",
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			path := r.URL.Path

			// Ensure path ends with .view (for Tempo)
			if !strings.HasSuffix(path, ".view") {
				path += ".view"
			}

			if path == rest.Fresh() {
				next.ServeHTTP(w, r)
				return
			}

			// 2. CHECK PATH VALIDITY
			isValidPath := slices.Contains(ValidPaths, path)
			isMockPath := slices.Contains(mockPaths, path)

			if !isValidPath && !isMockPath {
				w.WriteHeader(config.StatusNotFound)
				return
			}

			// 3. HANDLE MOCK ROUTES IMMEDIATELY
			if isMockPath {
				handleMockRoutes(w, r)
				return
			}

			if isValidPath && path != rest.Ping() && path != rest.GetUserView() {
				RewriteRequest(w, r)
				return
			}

			// 5. AUTHENTICATION & PING HANDLING
			// Only /rest/ping.view reaches here usually.

			/* Add authentication parameters
			to the URL query like -> (https://) */

			// salt := Salt(t)
			// token := Token(s, salt)

			// slog.Info("session details",
			// 	"user", userName,
			// 	"salt", salt,
			// 	"token", token,
			// )

			q := r.URL.Query()

			s := q.Get("s")
			t := q.Get("t")
			f := q.Get("f")
			c := q.Get("c")

			userName := q.Get("u")
			passWord := q.Get("p")

			if strings.HasPrefix(passWord, "enc:") {
				bp, err := hex.DecodeString(passWord[4:])
				if err == nil {
					passWord = string(bp)
				}
			}

			// Reconstruct query params for consistency
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
				"path", path,
				"user", userName,
			)

			// r.URL.Scheme = target.Scheme
			// r.URL.Host = target.Host
			// r.Host = target.Host

			// Mock Subsonic response for /ping endpoint

			// Admin bypass: skip authUrl entirely
			if config.MODE == "managed" {
				if userName == "hifi" && passWord == "local" {
					slog.Info("Skipping local auth server check")
					writeSubsonicv2(w, "ok", http.StatusOK, map[string]any{
						"user": map[string]any{"username": userName},
					})
					return
				}
			}

			authUrl := fmt.Sprintf("%s://%s/v1/secrets/?env=%s&path=/hifi_users&key=%s&app_id=%s", config.Http, config.ProxyHost, config.ENV, userName, config.AppID)

			client := &http.Client{}
			req, err := http.NewRequest("GET", authUrl, nil)

			if err != nil {
				slog.Error("failed to create auth request", "error", err)
				return
			}
			req.Header.Add("Authorization", "Bearer User "+config.ProxyKey)
			req.Header.Add(config.HeaderContentType, config.ContentTypeJSON)

			res, err := client.Do(req)
			if err != nil {
				slog.Error("failed to contact auth server", "error", err)
				http.Error(w, "Auth server unavailable", http.StatusBadGateway)
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				slog.Error("failed to read auth body", "error", err)
				return
			}

			var g types.AppGet

			if err := json.Unmarshal(body, &g); err != nil {
				slog.Error("error unmarshalling auth response", "error", err)
				return
			}

			// Check password
			match := true
			if len(g) > 0 {
				match, _ = argon2id.ComparePasswordAndHash(passWord, g[0].Password)
			}

			if !match {
				writeSubsonic(w, "failed", http.StatusBadRequest) // Subsonic expects 200 OK with failed body usually, or 401
			} else {

				if path == rest.GetUserView() {
					writeSubsonicv2(w, "ok", http.StatusOK, map[string]any{
						"user": map[string]any{"username": userName},
					})
				}
				if path == rest.Ping() {
					slog.Info("🏓 Ping request successful")
					writeSubsonic(w, "ok", http.StatusOK)
				}
			}
		})

		/* Forward the request to the
		subsonic server -> */

		// proxy.ServeHTTP(w, r)
	}
}

// Handle internal mock routes with standard responses
func handleMockRoutes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Ensure path ends with .view (for Tempo)
	if !strings.HasSuffix(path, ".view") {
		path += ".view"
	}

	switch path {
	case "/rest/getMusicFolders.view":
		slog.Info("🎭 Mocking getMusicFolders")
		writeSubsonicv2(w, "ok", http.StatusOK, map[string]any{
			"musicFolders": map[string]any{
				"musicFolder": []map[string]any{
					{"id": 1, "name": "Tidal/Hifi"},
				},
			},
		})

	case "/rest/getLicense.view":
		writeSubsonicv2(w, "ok", http.StatusOK, map[string]any{
			"license": map[string]any{
				"valid":          true,
				"email":          "valid@hifi.local",
				"licenseExpires": "2099-01-01T00:00:00",
			},
		})

	case "/rest/getPlaylists.view":
		slog.Info("🎭 Mocking getPlaylists (Empty)")
		writeSubsonicv2(w, "ok", http.StatusOK, map[string]any{
			"playlists": map[string]any{
				"playlist": []any{},
			},
		})

	case "/rest/getGenres.view":
		slog.Info("🎭 Mocking getGenres (Empty)")
		writeSubsonicv2(w, "ok", http.StatusOK, map[string]any{
			"genres": map[string]any{
				"genre": []any{},
			},
		})

	case "/rest/getStarred.view", "/rest/getStarred2.view":
		slog.Info("🎭 Mocking getStarred (Empty)")
		writeSubsonicv2(w, "ok", http.StatusOK, map[string]any{
			"starred": map[string]any{
				"song":   []any{},
				"album":  []any{},
				"artist": []any{},
			},
		})
	}
}
