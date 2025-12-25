package plugins

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/url"
	"strings"
	"time"

	"worker/config"
	"worker/types"

	"github.com/syumai/workers/cloudflare/fetch"
	"github.com/syumai/workers/cloudflare/kv"
)

const (
	kvNamespace   = "TIDAL"
	tokenKey      = "TIDAL_TOKEN"
	kvTTL         = 3600
	earlyRenewSec = 10
)

func TidalAuth(ctx context.Context) (string, error) {

	tokenKV, err := kv.NewNamespace(kvNamespace)
	if err != nil {
		return "", err
	}

	val, err := tokenKV.GetString(tokenKey, nil)
	if err == nil && val != "" {
		var data struct {
			Token      string `json:"token"`
			ExpiryUnix int64  `json:"expiry"`
		}
		if err := json.Unmarshal([]byte(val), &data); err == nil {
			now := time.Now().Unix()
			if now < data.ExpiryUnix-earlyRenewSec {
				slog.Info("💾 KV cache hit", "expiresInSec", data.ExpiryUnix-now)
				return data.Token, nil
			}
			slog.Info("⚠️ KV token near expiry, refreshing", "expiresInSec", data.ExpiryUnix-now)
		} else {
			slog.Warn("⚠️ Failed to unmarshal KV data, fetching new token", "error", err)
		}
	} else {
		slog.Info("❌ KV cache miss, fetching new token")
	}

	token, expiry, err := fetchToken(ctx)
	if err != nil {
		return "", err
	}

	data := struct {
		Token      string `json:"token"`
		ExpiryUnix int64  `json:"expiry"`
	}{
		Token:      token,
		ExpiryUnix: expiry,
	}
	bytes, _ := json.Marshal(data)
	if err := tokenKV.PutString(tokenKey, string(bytes), &kv.PutOptions{ExpirationTTL: kvTTL}); err != nil {
		slog.Warn("⚠️ Failed to write token to KV", "error", err)
	} else {
		slog.Info("✅ Token cached in KV", "expiresAt", time.Unix(expiry, 0).Local().Format("15:04:05"))
	}

	return token, nil
}

func fetchToken(ctx context.Context) (string, int64, error) {
	if config.ClientID == "" || config.ClientSecret == "" || config.RefreshToken == "" {
		return "", 0, errors.New("missing Tidal credentials")
	}

	form := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {config.RefreshToken},
	}

	apiURL := URLBuild(config.TidalAuthHost, "/v1/oauth2/token")
	req, err := fetch.NewRequest(ctx, "POST", apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", 0, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(config.ClientID + ":" + config.ClientSecret))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", config.ContentTypeForm)

	client := fetch.NewClient()
	res, err := client.Do(req, nil)
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		slog.Error("Tidal API error", "status", res.StatusCode, "body", string(body))
		return "", 0, errors.New("tidal API error")
	}

	var tokenResp types.TidalTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", 0, err
	}
	if tokenResp.AccessToken == "" {
		return "", 0, errors.New("empty access token")
	}

	expiry := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second).Unix()
	slog.Info("🪣 Fetched new token", "expiresAt", time.Unix(expiry, 0).Local().Format("15:04:05"))

	return tokenResp.AccessToken, expiry, nil
}
