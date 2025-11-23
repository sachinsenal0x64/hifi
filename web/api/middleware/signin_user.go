package middleware

import (
	"api/config"
	"api/types"
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func SigninUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req types.SigninRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	base := fmt.Sprintf("%s://%s", config.Scheme, config.ProxyHost)

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	loginCh := startLoginUser(ctx, client, base+"/rest/ping.view", req.Username, req.Password)

	res := <-loginCh

	if res.Err != nil {
		http.Error(w, res.Err.Error(), http.StatusBadGateway)
		return
	}

	if res.Status >= 400 {
		http.Error(w, string(res.Body), http.StatusBadRequest)
		return
	}

	var result types.RegisterResult

	result, err := registerUser(req.Username, req.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maxAge := 3600
	expiration := time.Now().Add(time.Second * time.Duration(maxAge))

	claims := &types.Claims{
		ID:       result.User.ID,
		Username: result.User.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		},
	}

	hash := sha256.Sum256([]byte(claims.RegisteredClaims.ID))
	tokenHashes[claims.RegisteredClaims.ID] = hash

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtSecret)

	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"token":  tokenString,
		"maxAge": fmt.Sprintf("%d", maxAge),
	})

}

func startLoginUser(ctx context.Context, client *http.Client, createURL, newUser, newPass string) <-chan types.CreateResult {
	out := make(chan types.CreateResult, 1)

	go func() {
		defer close(out)

		u, _ := url.Parse(createURL)
		q := u.Query()
		q.Set("u", newUser)
		q.Set("p", newPass)
		q.Set("f", "json")
		u.RawQuery = q.Encode()

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), nil)
		if err != nil {
			out <- types.CreateResult{Status: 0, Body: nil, Err: err}
			return
		}
		req.Header.Set(config.HeaderContentType, config.ContentTypeForm)

		resp, err := client.Do(req)
		if err != nil {
			out <- types.CreateResult{Status: 0, Body: nil, Err: err}
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var ping types.Ping
		if err := json.Unmarshal(body, &ping); err != nil {
			out <- types.CreateResult{Status: resp.StatusCode, Body: nil, Err: fmt.Errorf("failed to unmarshal JSON: %w", err)}
			return
		}

		if ping.SubsonicResponse.Status != "ok" {
			errMsg := fmt.Sprintf("failed to sign in user: %s", string(body))
			out <- types.CreateResult{
				Status: resp.StatusCode,
				Body:   []byte(ping.SubsonicResponse.Status),
				Err:    errors.New(errMsg),
			}
			return
		}

		out <- types.CreateResult{
			Status: resp.StatusCode,
			Body:   []byte(ping.SubsonicResponse.Status),
			Err:    nil,
		}

	}()
	return out
}
