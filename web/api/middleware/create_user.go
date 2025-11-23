package middleware

import (
	"api/config"
	"api/types"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
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

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	createURL := base + "/v1/secrets/?app_id=" + config.AppID + "&env=" + config.ENV

	createCh := startCreateUser(ctx, &http.Client{}, createURL, req.Username, req.Password)

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

		hash, err := argon2id.CreateHash(newPass, argon2id.DefaultParams)
		if err != nil {
			out <- types.CreateResult{Status: 0, Body: nil, Err: err}
		}

		payload := strings.NewReader(fmt.Sprintf(`{
  "secrets": [
    {
      "key": %q,
      "value": %q,
      "comment": "Hifi DB",
	  "path": "/hifi_users"

    }
  ]
}`, newUser, hash))

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, createURL, payload)

		if err != nil {
			out <- types.CreateResult{Status: 0, Body: nil, Err: err}
			return
		}

		req.Header.Add("Authorization", "Bearer User "+config.ProxyKey)
		req.Header.Add(config.HeaderContentType, config.ContentTypeJSON)

		resp, err := client.Do(req)

		if err != nil {
			out <- types.CreateResult{Status: 0, Body: nil, Err: err}
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var c types.AppCreate

		if err := json.Unmarshal(body, &c); err != nil {
			out <- types.CreateResult{Status: resp.StatusCode, Body: body, Err: nil}
			return
		}

		out <- types.CreateResult{Status: resp.StatusCode, Body: body, Err: nil}
	}()
	return out
}
