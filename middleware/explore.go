package middleware

import (
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/types"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

func fetchAndCacheAlbum(user, id string) types.SubsonicAlbum {

	album := make(map[string]string)

	albumMu.Lock()
	album[user] = id
	albumMu.Unlock()

	// fmt.Printf("[Fetch Album] %s\n", id)

	tidalURL := &url.URL{
		Scheme: config.Scheme,
		Host:   config.TidalHost,
		Path:   fmt.Sprintf("/v1/albums/%s/items", id),
	}
	q := tidalURL.Query()
	q.Set("countryCode", "US")
	q.Set("limit", "100") // Max limit = 100
	q.Set("offset", "0")
	tidalURL.RawQuery = q.Encode()

	req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)

	if config.MODE == "managed" {
		req.Header.Set("x-tidal-token:", config.ClientID)
	} else {
		req.Header.Set("Authorization", "Bearer "+TidalAuth())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("Failed to fetch album", "id", id, "error", err)
		return types.SubsonicAlbum{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read album body", "id", id, "error", err)
		return types.SubsonicAlbum{}
	}

	var tidalAlbum types.TidalAlbumResponse
	if err := json.Unmarshal(body, &tidalAlbum); err != nil {
		slog.Error("Failed to parse album JSON", "id", id, "error", err)
		return types.SubsonicAlbum{}
	}

	var albumResp types.SubsonicAlbum
	albumResp.ID = id
	albumResp.IsDir = true

	for _, item := range tidalAlbum.Items {
		song := types.SubsonicSong{
			ID:          fmt.Sprint(item.Item.ID),
			Duration:    item.Item.Duration,
			Title:       item.Item.Title,
			Album:       item.Item.Album.Title,
			Artist:      item.Item.Artist.Name,
			CoverArt:    item.Item.Album.Cover,
			Parent:      item.Item.Artist.ID,
			Type:        "music",
			IsVideo:     false,
			ContentType: "audio/flac",
			Suffix:      "flac",
			Year:        item.Item.StreamStartDate[0:4],
			Track:       item.Item.TrackNumber,
			ArtistID:    fmt.Sprint(item.Item.Artist.ID),
			AlbumID:     fmt.Sprint(item.Item.Album.ID),
		}

		albumResp.Parent = item.Item.Artist.ID
		albumResp.ArtistID = item.Item.Artist.ID
		albumResp.Name = item.Item.Album.Title
		albumResp.Title = item.Item.Album.Title
		albumResp.Artist = item.Item.Artist.Name
		albumResp.CoverArt = item.Item.Album.Cover
		albumResp.Year = item.Item.StreamStartDate[0:4]
		albumResp.SongCount = len(tidalAlbum.Items)
		albumResp.Duration += item.Item.Duration
		albumResp.Song = append(albumResp.Song, song)

	}

	useralbumMu.Lock()
	if userAlbumCache[user] == nil {
		userAlbumCache[user] = make(map[string]types.SubsonicAlbum)
	}
	userAlbumCache[user][id] = albumResp
	useralbumMu.Unlock()

	return albumResp
}
