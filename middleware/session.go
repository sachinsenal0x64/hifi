package middleware

import (
	"maps"
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

func setQueryParams(q url.Values, params map[string]string) {
	for k, v := range params {
		q.Set(k, v)
	}
}

// Helper to write standard Subsonic JSON responses for mocks
func writeMockResponse(w http.ResponseWriter, data any) {
	resp := map[string]any{
		"subsonic-response": map[string]any{
			"status":  "ok",
			"version": "1.16.1",
		},
	}

	// Merge the data into the inner map
	subResponse := resp["subsonic-response"].(map[string]any)

	// If data is a map, merge it. If it's nil, we just send status/version.
	if dataMap, ok := data.(map[string]any); ok {
		maps.Copy(subResponse, dataMap)
	}

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func Session(userName string, passWord string, ValidPaths []string) func(http.Handler) http.Handler {

	// Define internal mock paths that don't need the Rewrite logic but need valid JSON responses
	mockPaths := []string{
		"/rest/getUser.view",
		"/rest/getMusicFolders.view",
		"/rest/getLicense.view",
		"/rest/getPlaylists.view",
		"/rest/getGenres.view",
		"/rest/getStarred.view",
		"/rest/getStarred2.view",
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path == rest.Fresh() {
				next.ServeHTTP(w, r)
				return
			}

			// 2. CHECK PATH VALIDITY
			isValidPath := slices.Contains(ValidPaths, r.URL.Path)
			isMockPath := slices.Contains(mockPaths, r.URL.Path)

			if !isValidPath && !isMockPath {
				w.WriteHeader(config.StatusNotFound)
				return
			}

			// 3. HANDLE MOCK ROUTES IMMEDIATELY
			if isMockPath {
				handleMockRoutes(w, r)
				return
			}

			if isValidPath && r.URL.Path != rest.Ping() {
				RewriteRequest(w, r)
				return
			}

			// 5. AUTHENTICATION & PING HANDLING
			// Only /rest/ping.view reaches here usually.

			q := r.URL.Query()
			userName := q.Get("u")
			passWord := q.Get("p")
			s := q.Get("s")
			t := q.Get("t")
			f := q.Get("f")
			c := q.Get("c")

			// Reconstruct query params for consistency
			params := map[string]string{
				"u": userName,
				"c": c,
				"f": f,
			}

			if s != "" && t != "" {
				params["s"] = s
				params["t"] = t
			} else {
				params["p"] = passWord
			}

			setQueryParams(q, params)
			r.URL.RawQuery = q.Encode()

			slog.Info("incoming request",
				"path", r.URL.Path,
				"user", userName,
			)

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
			match := false
			if len(g) > 0 {
				match, _ = argon2id.ComparePasswordAndHash(passWord, g[0].Password)
			}

			if !match {
				writeSubsonic(w, "failed", http.StatusBadRequest) // Subsonic expects 200 OK with failed body usually, or 401
			} else {
				writeSubsonic(w, "ok", http.StatusOK)
			}
		})
	}
}

func handleMockRoutes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	user := r.URL.Query().Get("u")
	if user == "" {
		user = "admin"
	}

	switch path {
	case "/rest/getUser.view":
		slog.Info("🎭 Mocking getUser", "user", user)
		writeMockResponse(w, map[string]any{
			"user": map[string]any{
				"username":          user,
				"email":             "user@hifi.local",
				"scrobblingEnabled": true,
				"adminRole":         true,
				"settingsRole":      true,
				"downloadRole":      true,
				"streamRole":        true,
				"coverArtRole":      true,
				"shareRole":         true,
				"jukeboxRole":       false,
			},
		})

	case "/rest/getMusicFolders.view":
		slog.Info("🎭 Mocking getMusicFolders")
		writeMockResponse(w, map[string]any{
			"musicFolders": map[string]any{
				"musicFolder": []map[string]any{
					{"id": 1, "name": "Tidal/Hifi"},
				},
			},
		})

	case "/rest/getLicense.view":
		writeMockResponse(w, map[string]any{
			"license": map[string]any{
				"valid":          true,
				"email":          "valid@hifi.local",
				"licenseExpires": "2099-01-01T00:00:00",
			},
		})

	case "/rest/getPlaylists.view":
		slog.Info("🎭 Mocking getPlaylists (Empty)")
		writeMockResponse(w, map[string]any{
			"playlists": map[string]any{
				"playlist": []any{},
			},
		})

	case "/rest/getGenres.view":
		slog.Info("🎭 Mocking getGenres (Empty)")
		writeMockResponse(w, map[string]any{
			"genres": map[string]any{
				"genre": []any{},
			},
		})

	case "/rest/getStarred.view", "/rest/getStarred2.view":
		slog.Info("🎭 Mocking getStarred (Empty)")
		writeMockResponse(w, map[string]any{
			"starred": map[string]any{
				"song":   []any{},
				"album":  []any{},
				"artist": []any{},
			},
		})
	}
}
