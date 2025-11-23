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

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req types.UpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "missing authorization header", http.StatusUnauthorized)
		return
	}

	olduSername := r.Header.Get("X-Username")

	if olduSername == "" {
		http.Error(w, "missing username header", http.StatusUnauthorized)
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

	base := fmt.Sprintf("%s://%s", config.Scheme, config.ProxyHost)

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if req.Username == "" && req.Password == "" {
		http.Error(w, "no username or password provided", http.StatusBadRequest)
		return
	}

	if req.Username != "" && req.Password != "" {

		updateUsername := startUpdateUser(
			ctx, client,
			base+"/admin/change_username_do",
			olduSername,
			req.Username,
			startLogin(ctx, client, base+"/admin/login_do", config.ProxyKey),
		)

		resUsername := <-updateUsername

		if resUsername.Err != nil || resUsername.Status >= 400 {
			http.Error(w, "User update failed", http.StatusBadGateway)
			return
		}

		updatePassword := startUpdateUserPassword(
			ctx, client,
			base+"/admin/change_password_do",
			olduSername,
			req.Password,
			startLogin(ctx, client, base+"/admin/login_do", config.ProxyKey),
		)

		resPassword := <-updatePassword

		if resPassword.Err != nil || resPassword.Status >= 400 {
			http.Error(w, "User update failed", http.StatusBadGateway)
			return
		}

		mu.Lock()
		delete(users, olduSername)
		delete(tokenHashes, claims.RegisteredClaims.ID)

		user.Username = req.Username
		users[req.Username] = user
		mu.Unlock()

		w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "User updated successfully",
		})

		return
	}

	if req.Username != "" {

		updateUsername := startUpdateUser(ctx, client, base+"/admin/change_username_do", olduSername, req.Username, startLogin(ctx, client, base+"/admin/login_do", config.ProxyKey))
		res := <-updateUsername

		if res.Err != nil {
			http.Error(w, res.Err.Error(), http.StatusBadGateway)
			return
		}

		if res.Status >= 400 {
			http.Error(w, "User update failed", http.StatusBadGateway)
			return
		}

		mu.Lock()
		user.Username = req.Username
		users[req.Username] = user
		mu.Unlock()
	}
	if req.Password != "" {

		updatePassword := startUpdateUserPassword(ctx, client, base+"/admin/change_password_do", olduSername, req.Password, startLogin(ctx, client, base+"/admin/login_do", config.ProxyKey))
		res := <-updatePassword

		if res.Err != nil {
			http.Error(w, res.Err.Error(), http.StatusBadGateway)
			return
		}

		if res.Status >= 400 {
			http.Error(w, "User update failed", http.StatusBadGateway)
			return
		}

		mu.Lock()
		delete(users, claims.RegisteredClaims.ID)
		delete(tokenHashes, claims.RegisteredClaims.ID)
		mu.Unlock()
	}

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated successfully",
	})

}

func startUpdateUser(ctx context.Context, client *http.Client, updateURL, olduSername, reqUsername string, loginCh <-chan types.LoginResult) <-chan types.CreateResult {
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

		u, _ := url.Parse(updateURL)
		q := u.Query()
		q.Set("user", olduSername)
		u.RawQuery = q.Encode()

		form := url.Values{}
		form.Set("username", reqUsername)

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), strings.NewReader(form.Encode()))

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

func startUpdateUserPassword(ctx context.Context, client *http.Client, updateURL, olduSername, newPassword string, loginCh <-chan types.LoginResult) <-chan types.CreateResult {
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

		u, _ := url.Parse(updateURL)
		q := u.Query()
		q.Set("user", olduSername)
		u.RawQuery = q.Encode()

		form := url.Values{}
		form.Set("password_one", newPassword)
		form.Set("password_two", newPassword)

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), strings.NewReader(form.Encode()))

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
