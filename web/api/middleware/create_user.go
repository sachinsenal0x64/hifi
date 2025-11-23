package middleware

import (
	"api/config"
	"api/types"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
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

	base := fmt.Sprintf("%s://%s", config.Scheme, config.ProxyHost)

	slog.Info(base)

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	createCh := startCreateUser(ctx, &http.Client{}, base+"/v1/apps/secrets", req.Username, req.Password)

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

func startCreateUser(ctx context.Context, client *http.Client, createURL, newUser, newPass string) <-chan types.CreateResult {
	out := make(chan types.CreateResult, 1)
	go func() {
		defer close(out)

		base := fmt.Sprintf("%s://%s", config.Scheme, config.ProxyHost)

		form1 := url.Values{}
		form1.Set("name", newUser)
		form1.Set("scope[type]", "account")

		check, err := http.NewRequestWithContext(
			ctx, http.MethodGet,
			base+"/v1/apps/secrets/find",
			strings.NewReader(form1.Encode()),
		)

		if err != nil {
			out <- types.CreateResult{Status: 0, Body: nil, Err: err}
			return
		}

		checkResp, err := client.Do(check)
		if err != nil {
			out <- types.CreateResult{Status: 0, Body: nil, Err: err}
			return
		}

		defer checkResp.Body.Close()

		checkBody, _ := io.ReadAll(checkResp.Body)

		var f types.AppFind

		if err := json.Unmarshal(checkBody, &f); err != nil {
			out <- types.CreateResult{Status: checkResp.StatusCode, Body: checkBody, Err: err}
			return
		}

		if f.Name == newUser {
			out <- types.CreateResult{Status: checkResp.StatusCode, Body: checkBody,
				Err: fmt.Errorf("user already exists")}
			return
		}

		form := url.Values{}
		form.Set("name", newUser)
		form.Set("payload", newPass)
		form.Set("scope[type]", "account")

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, createURL, strings.NewReader(form.Encode()))

		if err != nil {
			out <- types.CreateResult{Status: 0, Body: nil, Err: err}
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			out <- types.CreateResult{Status: 0, Body: nil, Err: err}
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var c types.AppCreate

		if err := json.Unmarshal(checkBody, &c); err != nil {
			out <- types.CreateResult{Status: checkResp.StatusCode, Body: checkBody, Err: err}
			return
		}

		fmt.Println(c.Name)

		out <- types.CreateResult{Status: resp.StatusCode, Body: body, Err: nil}
	}()
	return out
}
