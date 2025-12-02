package middleware

import (
	"encoding/json"
	"hifi/config"
	"hifi/types"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/go-co-op/gocron/v2"
)

var (
	freshCache    []types.ExploreItem
	freshCacheMu  sync.RWMutex
	freshCacheExp time.Time
)

const (
	freshTTL          = 24 * time.Hour
	freshRefreshEvery = 23 * time.Hour
)

func refreshFreshCache() {
	items := GetNewAndTopItems()

	now := time.Now()
	exp := now.Add(freshTTL)

	freshCacheMu.Lock()
	freshCache = items
	freshCacheExp = exp
	freshCacheMu.Unlock()

	slog.Info("New TTL window started",
		"ttlMinutes", int(freshTTL),
		"expiresAt", exp.Format(time.RFC3339),
	)

	slog.Info("Fresh cache refreshed",
		"items", len(items),
		"refreshedAt", now.Format(time.RFC3339),
	)
}

func StartFreshRefresher() {
	s, err := gocron.NewScheduler()
	if err != nil {
		slog.Error("Failed to create fresh scheduler", "error", err)
	}

	refreshFreshCache()

	_, err = s.NewJob(
		gocron.DurationJob(freshRefreshEvery),
		gocron.NewTask(func() {

			slog.Info("Scheduled refresh triggered",
				"refreshEveryDays", int(freshRefreshEvery),
				"nextRefreshAt", time.Now().Add(freshRefreshEvery).Format(time.RFC3339),
			)

			refreshFreshCache()
		}),
	)
	if err != nil {
		slog.Error("Failed to schedule fresh refresher job", "error", err)
	}

	s.Start()

	slog.Info("Fresh cache refresher started",
		"refreshInterval", freshRefreshEvery.String(),
		"ttl", freshTTL.String())
}

func getFreshCachedItems() []types.ExploreItem {
	now := time.Now()

	freshCacheMu.RLock()
	valid := freshCache != nil && now.Before(freshCacheExp)
	items := freshCache

	freshCacheMu.RUnlock()

	if valid {
		return items
	}

	refreshFreshCache()

	freshCacheMu.RLock()
	defer freshCacheMu.RUnlock()
	return freshCache
}

func FreshHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	items := getFreshCachedItems()

	if len(items) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}

	var wg sync.WaitGroup

	for i := range items {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			items[idx].Cover = CoverFormat(items[idx].Cover)
		}(i)
	}

	wg.Wait()

	w.Header().Set(config.HeaderCacheControl, "public, max-age=86400")
	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)

	if err := json.NewEncoder(w).Encode(items); err != nil {
		slog.Error("failed to encode items", "error", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
