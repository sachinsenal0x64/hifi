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

func search3(search string, user string, w http.ResponseWriter) {

	var tidalSearch types.TidalSearchResponse

	if search != "" {
		queryMu.Lock()
		query[user] = search
		queryMu.Unlock()
	}

	queryMu.RLock()
	qu := query[user]
	queryMu.RUnlock()

	if qu == "" {
		qu = "sza essentials"
	}

	// Tidal search URL
	tidalURL := &url.URL{
		Scheme: config.Scheme,
		Host:   config.TidalHost,
		Path:   "/v2/search",
	}
	q := tidalURL.Query()
	q.Set("query", qu)
	q.Set("limit", "100") // Max limit = 100
	q.Set("offset", "0")
	q.Set("types", "ARTISTS,ALBUMS,TRACKS")
	q.Set("countryCode", "US")
	q.Set("deviceType", "BROWSER")
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

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Tidal returned %s", resp.Status), resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "failed to read Tidal response", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &tidalSearch); err != nil {
		http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
		return
	}

	sub := types.MetaBanner()
	sub.Subsonic.SearchResult3 = &types.SubsonicSearchResult{}

	// ARTISTS
	for _, a := range tidalSearch.Artists.Items {
		if a.Picture == "" {
			continue
		}
		artistID := a.ID

		artistsMu.RLock()
		userArtists := artistsCache[user]
		artistsMu.RUnlock()

		if userArtists != nil {
			if _, found := userArtists[artistID]; found {
				sub.Subsonic.SearchResult3.Artist = append(sub.Subsonic.SearchResult3.Artist, userArtists[artistID])
			}
		}

		artist := types.SubsonicArtist{
			ID:       fmt.Sprint(a.ID),
			Name:     a.Name,
			CoverArt: a.Picture,
		}

		artistsMu.Lock()
		if artistsCache[user] == nil {
			artistsCache[user] = make(map[int]types.SubsonicArtist)
		}
		if artistsOrder[user] == nil {
			artistsOrder[user] = []int{}
		}
		if _, exists := artistsCache[user][artistID]; !exists {
			artistsOrder[user] = append(artistsOrder[user], artistID)
		}
		artistsCache[user][artistID] = artist
		artistsMu.Unlock()

		sub.Subsonic.SearchResult3.Artist = append(sub.Subsonic.SearchResult3.Artist, artist)
	}

	// ALBUMS
	for _, alb := range tidalSearch.Albums.Items {

		albumID := fmt.Sprint(alb.ID)
		year := ""

		if len(alb.ReleaseDate) >= 4 {
			year = alb.ReleaseDate[:4]
		}

		albumYearMu.Lock()
		albumYearMap[albumID] = year
		albumYearMu.Unlock()

		sub.Subsonic.SearchResult3.Album = append(sub.Subsonic.SearchResult3.Album, types.SubsonicAlbum{
			ID:       albumID,
			Name:     alb.Title,
			Artist:   alb.Artist[0].Name,
			CoverArt: alb.Cover,
			Year:     year,
			IsDir:    true,
			Duration: alb.Duration,
		})
	}

	// SONGS
	for _, item := range tidalSearch.Tracks.Items {

		songID := fmt.Sprint(item.ID)

		songMu.RLock()
		userSongs := songMap[songID]
		songMu.RUnlock()

		if userSongs.ID != "" {
			slog.Info("Cached Hit", "id", songID)
			sub.Subsonic.SearchResult3.Song = append(sub.Subsonic.SearchResult3.Song, userSongs)
			continue
		}

		song := types.SubsonicSong{
			ID:          songID,
			Title:       item.Title,
			Artist:      item.Artist[0].Name,
			Album:       item.Album.Title,
			Duration:    item.Duration,
			CoverArt:    item.Album.Cover,
			Type:        "music",
			IsVideo:     false,
			ContentType: "audio/flac",
			Suffix:      "flac",
			ArtistID:    fmt.Sprint(item.Artist[0].ID),
			AlbumID:     fmt.Sprint(item.Album.ID),
		}

		sub.Subsonic.SearchResult3.Song = append(sub.Subsonic.SearchResult3.Song, song)

		songMu.Lock()
		songMap[songID] = song
		songMu.Unlock()

		slog.Info("Cached Miss", "id", songID)
	}

	// Write response
	json.NewEncoder(w).Encode(sub)

}
