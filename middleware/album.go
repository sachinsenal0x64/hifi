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

func getAlbum(id string, user string, w http.ResponseWriter) {

	var albumResp types.SubsonicAlbum
	var tidalAlbum types.TidalAlbumResponse

	album := make(map[string]string)

	albumMu.Lock()
	album[user] = id
	albumMu.Unlock()

	// Build Tidal request
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
		http.Error(w, fmt.Sprintf("tidal error: %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "failed to read Tidal response", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &tidalAlbum); err != nil {
		http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
		return
	}

	for _, item := range tidalAlbum.Items {

		songID := fmt.Sprint(item.Item.ID)

		song := types.SubsonicSong{
			ID:          songID,
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

		// Build Subsonic album
		albumID := fmt.Sprint(item.Item.Album.ID)

		albumYearMu.RLock()
		year, ok := albumYearMap[albumID]
		albumYearMu.RUnlock()

		if ok && year != "" {
			albumResp.Year = year
		} else {
			albumResp.Year = item.Item.StreamStartDate[:4]
		}

		albumResp.ID = albumID
		albumResp.IsDir = true
		albumResp.Parent = item.Item.Artist.ID
		albumResp.ArtistID = item.Item.Artist.ID
		albumResp.Name = item.Item.Album.Title
		albumResp.Title = item.Item.Album.Title
		albumResp.Artist = item.Item.Artist.Name
		albumResp.CoverArt = item.Item.Album.Cover
		albumResp.SongCount = tidalAlbum.TotalNumberOfItems
		albumResp.Duration += item.Item.Duration

		songMu.RLock()
		userSongs := songMap[songID]
		songMu.RUnlock()

		if userSongs.ID != "" {
			slog.Info("Cached Hit", "id", songID)
			albumResp.Song = append(albumResp.Song, userSongs)
			continue
		}

		songMu.Lock()
		songMap[songID] = song
		songMu.Unlock()
		slog.Info("Cached Miss", "id", songID)
		albumResp.Song = append(albumResp.Song, song)

	}

	useralbumMu.Lock()
	if userAlbumCache[user] == nil {
		userAlbumCache[user] = make(map[string]types.SubsonicAlbum)
	}
	userAlbumCache[user][id] = albumResp
	useralbumMu.Unlock()

	sub := types.MetaBanner()

	sub.Subsonic.Album = &albumResp
	json.NewEncoder(w).Encode(sub)

}
