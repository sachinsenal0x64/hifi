package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/routes/rest"
	"hifi/types"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// -------------------- REWRITE --------------------

var (
	songMap  = make(map[string]types.SubsonicSong)
	coverMap = make(map[string]string)
)

type PlaybackInfo struct {
	Manifest string `json:"manifest"`
}

type ManifestData struct {
	Urls []string `json:"urls"`
}

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

		req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
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
		sub.Subsonic.SearchResult3 = &types.SubsonicSearchResult{}

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
			}

			// Song
			song := types.SubsonicSong{
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
			}

			sub.Subsonic.SearchResult3.Song = append(sub.Subsonic.SearchResult3.Song, song)
			songMap[song.ID] = song

		}

		// Write response
		w.Header().Set("Cache-Control", "public, max-age=3600, immutable")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sub)

	// -------------------- COVER ART --------------------

	case rest.GetCoverArtView():
		id := r.URL.Query().Get("id")
		size := r.URL.Query().Get("size")

		uuid := id

		sizeMapping := map[int]int{
			100: 640,
			200: 640,
			300: 80,
			450: 640,
		}

		s, _ := strconv.Atoi(size)
		mappedSize := sizeMapping[s]

		redirectURL := fmt.Sprintf(
			"https://%s/images/%s/%dx%d.jpg",
			config.TidalStaticHost,
			FormatCoverID(uuid),
			mappedSize,
			mappedSize,
		)

		http.Redirect(w, r, redirectURL, config.StatusRedirectPermanent)

	// -------------------- getSong --------------------

	case rest.GetSong():
		id := r.URL.Query().Get("id")

		song := songMap[id]

		sub := types.MetaBanner()
		sub.Subsonic.Song = &song

		w.Header().Set("Cache-Control", "public, max-age=3600, immutable")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sub)

	// -------------------- Scrobble --------------------

	case rest.Scrobble():
		return

	// -------------------- Stream --------------------

	case rest.Stream():
		id := r.URL.Query().Get("id")

		tidalURL := &url.URL{
			Scheme: config.Scheme,
			Host:   config.TidalHost,
			Path:   fmt.Sprintf("/v1/tracks/%s/playbackinfopostpaywall/v4", id),
		}

		q := tidalURL.Query()
		q.Set("audioquality", "LOSSLESS")
		q.Set("playbackmode", "STREAM")
		q.Set("assetpresentation", "FULL")
		tidalURL.RawQuery = q.Encode()

		req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
		req.Header.Set("Authorization", "Bearer "+TidalAuth())

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "failed to request Tidal API", config.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		var playback PlaybackInfo
		if err := json.Unmarshal(body, &playback); err != nil {
			http.Error(w, "failed to parse playback info", config.StatusInternalServerError)
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(playback.Manifest)

		if err != nil {
			http.Error(w, "failed to decode manifest", config.StatusInternalServerError)
			return
		}

		var manifest ManifestData
		if err := json.Unmarshal(decoded, &manifest); err != nil {
			http.Error(w, "failed to parse manifest JSON", config.StatusInternalServerError)
			return
		}

		if len(manifest.Urls) == 0 {
			http.Error(w, "no stream URLs found", config.StatusNotFound)
			return
		}

		streamURL := manifest.Urls[0]
		http.Redirect(w, r, streamURL, config.StatusRedirectPermanent)

	}
}
