package middleware

import (
	"api/config"
	"api/types"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func SignupUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req types.SignupRequest

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

	base := fmt.Sprintf("%s://%s", config.HifiScheme, config.ProxyHost)

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	createCh := startCreateUser(ctx, client, base+"/v1/apps/secrets", req.Username, req.Password, startLogin(ctx, client, base+"/admin/login_do", config.ProxyKey))

	res := <-createCh

	if res.Err != nil {
		http.Error(w, res.Err.Error(), http.StatusBadGateway)
		return
	}

	if res.Status >= 400 {
		http.Error(w, "User creation failed", http.StatusBadGateway)
		return
	}

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func startCreateUser(ctx context.Context, client *http.Client, createURL, newUser, newPass string, loginCh <-chan types.LoginResult) <-chan types.CreateResult {
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

		form := url.Values{}
		form.Set("username", newUser)
		form.Set("password_one", newPass)
		form.Set("password_two", newPass)

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, createURL, strings.NewReader(form.Encode()))

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
