package middleware

import (
	"api/config"
	"api/types"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req types.DeleteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	if req.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "missing authorization header", http.StatusUnauthorized)
		return
	}

	var tokenString string
	fmt.Sscanf(authHeader, "Bearer %s", &tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &types.Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return config.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(*types.Claims)
	if !ok {
		http.Error(w, "Invalid claims", http.StatusBadRequest)
		return
	}

	computed := sha256.Sum256([]byte(claims.RegisteredClaims.ID))

	stored, ok := tokenHashes[claims.RegisteredClaims.ID]
	if !ok || !bytes.Equal(stored[:], computed[:]) {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	mu.RLock()
	user, exists := users[claims.Username]
	mu.RUnlock()

	if !exists || user.ID != claims.ID {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}

	if !exists || req.Username != claims.Username {
		http.Error(w, "User does not exist or data mismatch", http.StatusBadRequest)
		return
	}

	mu.Lock()
	delete(users, req.Username)
	delete(users, claims.Username)
	delete(tokenHashes, claims.RegisteredClaims.ID)
	mu.Unlock()

	base := fmt.Sprintf("%s://%s", config.HifiScheme, config.ProxyHost)

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	createCh := startDeleteUser(ctx, client, base+"/admin/delete_user_do", req.Username, startLogin(ctx, client, base+"/admin/login_do", config.ProxyKey))

	res := <-createCh

	if res.Err != nil {
		http.Error(w, res.Err.Error(), http.StatusBadGateway)
		return
	}

	if res.Status >= 400 {
		http.Error(w, "User deletion failed", http.StatusBadGateway)
		return
	}

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

func startDeleteUser(ctx context.Context, client *http.Client, deleteURL, username string, loginCh <-chan types.LoginResult) <-chan types.CreateResult {
	out := make(chan types.CreateResult, 1)
	go func() {
		defer close(out)

		select {
		case lr := <-loginCh:
			if lr.Err != nil || !lr.OK {
				out <- types.CreateResult{Err: fmt.Errorf("login failed")}
				return
			}

		case <-ctx.Done():
			out <- types.CreateResult{Status: 0, Body: nil, Err: ctx.Err()}
			return
		}

		u, _ := url.Parse(deleteURL)
		q := u.Query()
		q.Set("user", username)
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

		out <- types.CreateResult{Status: resp.StatusCode, Body: body, Err: nil}
	}()
	return out
}
