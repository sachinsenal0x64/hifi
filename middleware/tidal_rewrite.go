package middleware

import (
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/routes/rest"
	"hifi/types"
	"io"
	"net/http"
	"net/url"
)

// -------------------- REWRITE --------------------

var coverMap = make(map[string]string)

func RewriteRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case rest.Search3View():
		query := r.URL.Query().Get("query")

		// Tidal search URL
		tidalURL := &url.URL{
			Scheme: config.Scheme,
			Host:   config.TidalHost,
			Path:   "/v1/search/tracks",
		}
		q := tidalURL.Query()
		q.Set("query", query)
		q.Set("limit", "1500")
		q.Set("offset", "0")
		q.Set("countryCode", "US")
		tidalURL.RawQuery = q.Encode()

		req, _ := http.NewRequest(http.MethodGet, tidalURL.String(), nil)
		req.Header.Set("Authorization", "Bearer "+TidalAuth())

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

		var tidal types.TidalResponse
		if err := json.Unmarshal(body, &tidal); err != nil {
			http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
			return
		}

		sub := types.MetaBanner()

		if sub.Subsonic.SearchResult3 == nil {
			sub.Subsonic.SearchResult3 = &types.SubsonicSearchResult{}
		}

		artistMap := make(map[int]bool)
		albumMap := make(map[int]bool)

		for _, item := range tidal.Items {

			albumID := fmt.Sprint(item.Album.ID)
			songID := fmt.Sprint(item.ID)
			coverUUID := item.Album.Cover

			coverMap[albumID] = coverUUID // album
			coverMap[songID] = coverUUID  // song

			// Artist
			if !artistMap[item.Artist.ID] {
				sub.Subsonic.SearchResult3.Artist = append(sub.Subsonic.SearchResult3.Artist, types.SubsonicArtist{
					ID:         fmt.Sprint(item.Artist.ID),
					Name:       item.Artist.Name,
					CoverArt:   item.Artist.Picture,
					AlbumCount: 11,
				})
				artistMap[item.Artist.ID] = true
				fmt.Println("Added artist:", item.Artist.Picture)
			}

			// Album
			if !albumMap[item.Album.ID] {
				sub.Subsonic.SearchResult3.Album = append(sub.Subsonic.SearchResult3.Album, types.SubsonicAlbum{
					ID:        fmt.Sprint(item.Album.ID),
					Name:      item.Album.Title,
					Artist:    item.Artist.Name,
					CoverArt:  item.Album.Cover,
					SongCount: 11,
				})
				albumMap[item.Album.ID] = true
				fmt.Println("Added album:", item.Album.Cover)
			}

			// Song
			sub.Subsonic.SearchResult3.Song = append(sub.Subsonic.SearchResult3.Song, types.SubsonicSong{
				ID:          fmt.Sprint(item.ID),
				Title:       item.Title,
				Album:       item.Album.Title,
				Artist:      item.Artist.Name,
				Duration:    item.Duration,
				CoverArt:    item.Album.Cover,
				Type:        "music",
				IsVideo:     false,
				ContentType: "audio/flac",
				Suffix:      "flac",
				ArtistID:    fmt.Sprint(item.Artist.ID),
				AlbumID:     fmt.Sprint(item.Album.ID),
			})

			fmt.Println("Added album:", item.ID)

		}

		// Write response
		w.Header().Set("Cache-Control", "public, max-age=3600, immutable")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sub)

		// -------------------- COVER ART --------------------

	case rest.GetCoverArtView():
		id := r.URL.Query().Get("id")

		uuid := id
		if v, ok := coverMap[id]; ok {
			uuid = v
		}

		path := FormatCoverID(uuid)
		redirectURL := fmt.Sprintf("https://%s/images/%s/80x80.jpg", config.TidalStaticHost, path)

		http.Redirect(w, r, redirectURL, config.StatusRedirectPermanent)

	// -------------------- MOCKS --------------------

	case rest.GetArtistsView():
		sub := types.MetaBanner()

		if sub.Subsonic.Artists == nil {
			sub.Subsonic.Artists = &types.SubsonicArtists{}
		}

		sub.Subsonic.Artists.IgnoredArticles = "The An A Die Das Ein Eine Les Le La"

		sub.Subsonic.Artists.Index = []types.SubsonicArtistIndexItem{
			{
				Name: "C",
				Artist: []types.SubsonicArtist{
					{ID: "100000016", Name: "CARNÚN", CoverArt: "ar-100000016", AlbumCount: 1},
					{ID: "100000027", Name: "Chi.Otic", CoverArt: "ar-100000027", AlbumCount: 0},
				},
			},
			{
				Name: "I",
				Artist: []types.SubsonicArtist{
					{ID: "100000013", Name: "IOK-1", CoverArt: "ar-100000013", AlbumCount: 1},
				},
			},
		}

		w.Header().Set("Cache-Control", "public, max-age=3600, immutable")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sub)

	}
}
