package plugins

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/url"
	"strconv"
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
	expiryKey     = "TIDAL_TOKEN_EXPIRY"
	earlyRenewSec = 10
)

func TidalAuth(ctx context.Context) (string, error) {
	tokenKV, err := kv.NewNamespace(kvNamespace)
	if err != nil {
		return "", err
	}

	// Try fetching token and expiry from KV
	token, _ := tokenKV.GetString(tokenKey, nil)
	expiryStr, _ := tokenKV.GetString(expiryKey, nil)

	if token != "" && expiryStr != "" {
		expiryUnix, err := strconv.ParseInt(expiryStr, 10, 64)
		if err == nil && time.Now().Unix() < expiryUnix-earlyRenewSec {
			slog.Info("💾 KV cache hit", "expiresInSec", expiryUnix-time.Now().Unix())
			return token, nil
		}
		slog.Info("⚠️ KV token expired or near expiry, refreshing")
	} else {
		slog.Info("❌ KV cache miss, fetching new token")
	}

	// Fetch new token
	newToken, expiry, err := fetchToken(ctx)
	if err != nil {
		return "", err
	}

	// Store token and expiry separately in KV
	ttl := int(expiry - time.Now().Unix() - earlyRenewSec)
	tokenKV.PutString(tokenKey, newToken, &kv.PutOptions{ExpirationTTL: ttl})
	tokenKV.PutString(expiryKey, strconv.FormatInt(expiry, 10), &kv.PutOptions{ExpirationTTL: ttl})

	slog.Info("✅ Token and expiry cached in KV", "expiresAt", time.Unix(expiry, 0).Local().Format("15:04:05"))

	return newToken, nil
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
