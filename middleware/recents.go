package middleware

import (
	"hifi/types"
	"log/slog"
	"strconv"
)

func RecentAlbum() {

	var allAlbums []types.SubsonicAlbum

	ids := getFreshCachedItems()

	if len(ids) == 0 {
		slog.Warn("No IDs returned from GetNewAndTop")
		return
	}

	user := Public
	results := make(chan types.SubsonicAlbum, len(ids))

	for _, id := range ids {
		go func(albumID string) {
			album := fetchAndCacheAlbum(user, albumID)
			results <- album
		}(strconv.Itoa(id.ID))
	}

	for range ids {
		album := <-results
		// fmt.Printf("[Album cached] %s — %s\n", album.ID, album.Title)
		allAlbums = append(allAlbums, album)
	}
	close(results)

	useralbumMu.Lock()
	if userAlbumCache[user] == nil {
		userAlbumCache[user] = make(map[string]types.SubsonicAlbum)
	}
	for _, album := range allAlbums {
		userAlbumCache[user][album.ID] = album
	}
	useralbumMu.Unlock()

	logger := slog.Default()
	logger.Info("Album refresh completed",
		slog.Int("total_albums_cached", len(allAlbums)),
		slog.String("user", user),
	)
}
