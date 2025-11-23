package middleware

import (
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/types"
	"io"
	"net/http"
	"net/url"
)

func getArtist(id string, user string, w http.ResponseWriter) {
	var tidalArtistAlbums types.TidalArtistAlbumsResponse

	artistID := id

	sub := types.MetaBanner()
	sub.Subsonic.Artist = &types.SubsonicArtistWithAlbums{}

	artistMu.RLock()
	if perUser := artistsWithAlbumsCache[user]; perUser != nil {
		if cached, ok := perUser[artistID]; ok {
			sub.Subsonic.Artist = &cached
			artistMu.RUnlock()
			_ = json.NewEncoder(w).Encode(sub)
			return
		}
	}
	artistMu.RUnlock()

	tidalURL := &url.URL{
		Scheme: config.Scheme,
		Host:   config.TidalHost,
		Path:   fmt.Sprintf("/v2/artist/%s", id),
	}
	q := tidalURL.Query()
	q.Set("locale", "en_US")
	q.Set("countryCode", "US")
	q.Set("deviceType", "BROWSER")
	q.Set("platform", "WEB")
	tidalURL.RawQuery = q.Encode()

	req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
	req.Header.Set("Authorization", "Bearer "+TidalAuth())
	req.Header.Set("x-tidal-client-version", "2025.11.3")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Tidal error: %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read Tidal response", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &tidalArtistAlbums); err != nil {
		http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
		return
	}

	var albums []types.SubsonicAlbum

	validModules := map[string]bool{
		"ARTIST_ALBUMS":      true,
		"ARTIST_TOP_SINGLES": true,
		"ARTIST_LIVE_ALBUMS": true,
		"ARTIST_APPEARS_ON":  true,
	}

	for _, item := range tidalArtistAlbums.Items {
		if validModules[item.ModuleId] {
			for _, albumItem := range item.Items {
				data := albumItem.Data
				id := fmt.Sprint(data.ID)
				title := data.Title
				duration := data.Duration
				year := data.ReleaseDate[:4]
				cover := data.Cover

				albums = append(albums, types.SubsonicAlbum{
					ID:       id,
					Name:     title,
					Year:     year,
					Duration: duration,
					CoverArt: cover,
				})
			}
		}
	}

	artistData := tidalArtistAlbums.Item.Data

	artist := types.SubsonicArtistWithAlbums{
		ID:         fmt.Sprint(artistData.ID),
		Name:       artistData.Name,
		CoverArt:   firstNonEmpty(artistData.Picture, artistData.SelectedAlbumCoverFallback),
		AlbumCount: len(albums),
		Album:      albums,
	}

	artistMu.Lock()
	if artistsWithAlbumsCache[user] == nil {
		artistsWithAlbumsCache[user] = make(map[string]types.SubsonicArtistWithAlbums)
	}
	artistsWithAlbumsCache[user][artistID] = artist
	artistMu.Unlock()

	sub.Subsonic.Artist = &artist

	_ = json.NewEncoder(w).Encode(sub)
}
