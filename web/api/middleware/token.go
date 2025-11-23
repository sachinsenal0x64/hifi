package middleware

import (
	"api/config"
	"api/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func startLoginAdminUser(ctx context.Context, client *http.Client, createURL, newPass string) <-chan types.CreateResult {
	out := make(chan types.CreateResult, 1)

	go func() {
		defer close(out)

		u, _ := url.Parse(createURL)
		q := u.Query()
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

func startLogin(ctx context.Context, client *http.Client, loginDoURL, pass string) <-chan types.LoginResult {
	token := make(chan types.LoginResult, 1)

	go func() {
		defer close(token)

		form := url.Values{}
		form.Set("password", pass)

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, loginDoURL, strings.NewReader(form.Encode()))
		if err != nil {
			token <- types.LoginResult{OK: false, Err: err}
			return
		}
		req.Header.Set(config.HeaderContentType, config.ContentTypeForm)

		resp, err := client.Do(req)

		if err != nil {
			token <- types.LoginResult{OK: false, Err: err}
			return
		}
		defer resp.Body.Close()

		base := fmt.Sprintf("%s://%s", config.Scheme, config.ProxyHost)

		login := startLoginAdminUser(ctx, client, base+"/rest/ping.view", pass)

		res := <-login

		if res.Err != nil {
			token <- types.LoginResult{OK: false, Err: fmt.Errorf("invalid login: %d", res.Status)}
			return
		}

		if res.Status >= 400 {
			token <- types.LoginResult{OK: false, Err: fmt.Errorf("login failed: %d", res.Status)}
			return
		}

		token <- types.LoginResult{OK: true, Err: nil}
	}()
	return token
}
