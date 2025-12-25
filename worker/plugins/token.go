package plugins

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
	"worker/config"
	"worker/types"

	"github.com/syumai/workers/cloudflare/fetch"

	"github.com/go-co-op/gocron/v2"
)

var (
	mu              sync.Mutex
	cachedToken     string
	TokenExpiryUnix int64
	scheduler       gocron.Scheduler
)

const earlyWarningSeconds = 10

func StartTidalRefresher() {
	s, err := gocron.NewScheduler()
	if err != nil {
		slog.Error("Failed to create scheduler", "error", err)
		return
	}

	mu.Lock()
	scheduler = s
	mu.Unlock()

	s.Start()
	slog.Info("Token refresher started")

	// Initial token fetch
	refreshAndReschedule("startup")

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	slog.Info("🛑 Shutting down Tidal refresher...")

	mu.Lock()
	if scheduler != nil {
		scheduler.Shutdown()
	}
	mu.Unlock()
}

// --- Scheduler Setup ---
func setupTokenJobs() {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now().Unix()
	remaining := TokenExpiryUnix - now
	s := scheduler

	// Remove existing jobs to prevent duplication
	if s != nil {
		for _, job := range s.Jobs() {
			_ = s.RemoveJob(job.ID())
		}
	}

	if remaining <= 0 {
		slog.Warn("Token expired — refreshing now.")
		go refreshAndReschedule("expired-immediate")
		return
	}

	slog.Info("📅 Scheduling token jobs",
		"expiresInSec", remaining,
		"earlyWarningSec", earlyWarningSeconds,
		"expiresAt", time.Unix(TokenExpiryUnix, 0).Local().Format("Mon, 02 Jan 2006 15:04:05 GMT-07"))

	// --- Early refresh ---
	if remaining > earlyWarningSeconds && s != nil {
		delay := remaining - earlyWarningSeconds
		_, err := s.NewJob(
			gocron.DurationJob(time.Duration(delay)*time.Second),
			gocron.NewTask(func() {
				slog.Warn("⚠️ Token expiring soon", "secondsLeft", earlyWarningSeconds)
				refreshAndReschedule("early-warning")
			}),
		)
		if err != nil {
			slog.Error("📛 Failed to schedule early warning job", "error", err)
		}
	}
}

// --- Refresh Token and Reschedule ---
func refreshAndReschedule(reason string) {
	// slog.Info("🔁 Refreshing Tidal token...", "reason", reason)

	token, err := refreshToken()
	if err != nil {
		slog.Error("📛 Token refresh failed", "reason", reason, "error", err)
		time.AfterFunc(30*time.Second, func() { refreshAndReschedule(reason) })
		return
	}

	mu.Lock()
	expiry := TokenExpiryUnix
	mu.Unlock()

	slog.Info("✅ Token refreshed successfully",
		"reason", reason,
		"token", token[:10]+"...",
		"expiresAt", time.Unix(expiry, 0).Local().Format("Mon, 02 Jan 2006 15:04:05 GMT-07"))

	setupTokenJobs()
}

// --- API Request for New Token ---
func refreshToken() (string, error) {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now().Unix()
	if cachedToken != "" && now < TokenExpiryUnix {
		// slog.Info("💾 Token cache hit",
		// 	"expiresAt", time.Unix(TokenExpiryUnix, 0).Local().Format("15:04:05"),
		// 	"secondsLeft", TokenExpiryUnix-now)
		return cachedToken, nil
	}

	slog.Warn("🆕 Token cache miss — requesting new token from Tidal API")

	if config.ClientID == "" || config.ClientSecret == "" || config.RefreshToken == "" {
		return "", errors.New("missing Tidal credentials")
	}

	form := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {config.RefreshToken},
	}

	apiURL := URLBuild(config.TidalAuthHost, "/v1/oauth2/token")
	req, err := fetch.NewRequest(context.Background(), http.MethodPost, apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(config.ClientID + ":" + config.ClientSecret))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", config.ContentTypeForm)

	cli := fetch.NewClient()

	res, err := cli.Do(req, nil)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		slog.Error("Tidal API error", "status", res.StatusCode, "body", string(body))
		return "", errors.New("tidal API error")
	}

	var token types.TidalTokenResponse
	if err := json.Unmarshal(body, &token); err != nil {
		return "", err
	}
	if token.AccessToken == "" {
		return "", errors.New("empty access token in response")
	}

	cachedToken = token.AccessToken

	fmt.Println(cachedToken)

	TokenExpiryUnix = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second).Unix()

	slog.Info("🪣 Cached new token", "expiresAt", time.Unix(TokenExpiryUnix, 0).Local().Format("15:04:05"))

	return cachedToken, nil
}

func TidalAuth() string {
	token, err := refreshToken()
	if err != nil {
		slog.Error("Failed to get Tidal token", "error", err)
		return ""
	}
	return token
}
