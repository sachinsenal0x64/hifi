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
	"sync"
	"time"

	"worker/config"
	"worker/types"

	"github.com/syumai/workers/cloudflare/fetch"
)

var (
	mu          sync.Mutex
	cachedToken string
	tokenExpiry int64
)

const earlyRenewSec = 10

func TidalAuth(ctx context.Context) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now().Unix()

	// Cache hit
	if cachedToken != "" && now < tokenExpiry-earlyRenewSec {
		slog.Info("💾 Cache hit", "expiresInSec", tokenExpiry-now)
		return cachedToken, nil
	}

	if cachedToken != "" {
		slog.Info("⚠️ Token near expiry, refreshing")
	} else {
		slog.Info("❌ Cache miss, fetching new token")
	}

	token, expiry, err := fetchToken(ctx)
	if err != nil {
		return "", err
	}

	cachedToken = token
	tokenExpiry = expiry
	return cachedToken, nil
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
