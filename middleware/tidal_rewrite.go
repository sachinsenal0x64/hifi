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
	coverMap = make(map[string]string)
	songMap  = make(map[string]types.SubsonicSong)
)

func GetPlaybackInfo(trackID string) (string, error) {

	// Tidal playback URL
	playbackURL := &url.URL{
		Scheme: config.Scheme,
		Host:   config.TidalHost,
		Path:   fmt.Sprintf("/v1/%s/playbackinfopostpaywall/v4", trackID),
	}

	q := playbackURL.Query()
	q.Set("audioquality", "LOSSLESS")
	q.Set("playbackmode", "STREAM")
	q.Set("assetpresentation", "FULL")
	playbackURL.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, playbackURL.String(), nil)
	req.Header.Set("Authorization", "Bearer "+TidalAuth())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var playback struct {
		Manifest string `json:"manifest"`
	}
	if err := json.Unmarshal(body, &playback); err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(playback.Manifest)
	if err != nil {
		return "", err
	}

	var manifest struct {
		Urls []string `json:"urls"`
	}
	if err := json.Unmarshal(decoded, &manifest); err != nil {
		return "", err
	}

	if len(manifest.Urls) == 0 {
		return "", fmt.Errorf("no urls in manifest")
	}

	return manifest.Urls[0], nil
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
		sub.Subsonic.SearchResult3 = &types.SubsonicSearchResult{}

		for _, item := range tidal.Items {

			albumID := fmt.Sprint(item.Album.ID)
			songID := fmt.Sprint(item.ID)
			coverUUID := item.Album.Cover

			coverMap[albumID] = coverUUID // album
			coverMap[songID] = coverUUID  // song

			// Artist
			sub.Subsonic.SearchResult3.Artist = append(sub.Subsonic.SearchResult3.Artist, types.SubsonicArtist{
				ID:         fmt.Sprint(item.Artist.ID),
				Name:       item.Artist.Name,
				CoverArt:   item.Artist.Picture,
				AlbumCount: 11,
			})

			// Album
			sub.Subsonic.SearchResult3.Album = append(sub.Subsonic.SearchResult3.Album, types.SubsonicAlbum{
				ID:        fmt.Sprint(item.Album.ID),
				Name:      item.Album.Title,
				Artist:    item.Artist.Name,
				CoverArt:  item.Album.Cover,
				SongCount: 11,
			})

			// Song
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
			80:  80,
			300: 80,
			450: 320,
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
		playbackURL, err := GetPlaybackInfo(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("playback error: %v", err), http.StatusBadGateway)
			return
		}

		http.Redirect(w, r, playbackURL, config.StatusRedirectPermanent)

	}
}
