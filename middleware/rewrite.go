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
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// -------------------- REWRITE --------------------

var (
	songMap     = make(map[string]types.SubsonicSong)
	coverMap    = make(map[string]string)
	playback    types.PlaybackInfo
	manifest    types.ManifestData
	tidalSearch types.TidalSearchResponse
)

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

		if err := json.Unmarshal(body, &tidalSearch); err != nil {
			http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
			return
		}

		sub := types.MetaBanner()
		sub.Subsonic.SearchResult3 = &types.SubsonicSearchResult{}

		artistMap := make(map[int]bool)
		albumMap := make(map[int]bool)

		for _, item := range tidalSearch.Items {

			albumID := fmt.Sprint(item.Album.ID)
			songID := fmt.Sprint(item.ID)
			coverUUID := item.Album.Cover

			coverMap[albumID] = coverUUID // album
			coverMap[songID] = coverUUID  // song

			// Artist
			if !artistMap[item.Artist.ID] {
				sub.Subsonic.SearchResult3.Artist = append(sub.Subsonic.SearchResult3.Artist, types.SubsonicArtist{
					ID:       fmt.Sprint(item.Artist.ID),
					Name:     item.Artist.Name,
					CoverArt: item.Artist.Picture,
				})
				artistMap[item.Artist.ID] = true
			}

			// Album
			if !albumMap[item.Album.ID] {
				sub.Subsonic.SearchResult3.Album = append(sub.Subsonic.SearchResult3.Album, types.SubsonicAlbum{
					ID:       fmt.Sprint(item.Album.ID),
					Name:     item.Album.Title,
					Artist:   item.Artist.Name,
					CoverArt: item.Album.Cover,
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
			100: 160,
			200: 320,
			300: 320,
			450: 640, // 1080x1080
			500: 750, // 1280x1280
		}

		s, _ := strconv.Atoi(size)
		mappedSize := sizeMapping[s]

		redirectURL := fmt.Sprintf(
			"%s://%s/images/%s/%dx%d.jpg",
			config.Scheme,
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

	// -------------------- getArtists --------------------

	case rest.GetArtistsView():

		songMap := []types.TidalArtistResponse{
			{ArtistID: 1, Name: "Nirvana", ProfilePicture: "cover1.jpg"},
			{ArtistID: 2, Name: "A Radiohead", ProfilePicture: "cover2.jpg"},
			{ArtistID: 2, Name: "jack", ProfilePicture: "cover1.jpg"}, // duplicate test
			{ArtistID: 4, Name: "Swift", ProfilePicture: "0af461b8/96b3/4711/a27b/1cb765263111"},
			{ArtistID: 5, Name: "Coldplay", ProfilePicture: "cover3.jpg"},
			{ArtistID: 6, Name: "The Beatles", ProfilePicture: "cover4.jpg"},
			{ArtistID: 7, Name: "An Endless Sporadic", ProfilePicture: "cover5.jpg"},
		}

		sub := types.MetaBanner()
		sub.Subsonic.Artists = &types.SubsonicArtists{}

		// Define ignored articles
		ignoredArticles := []string{"The", "An", "A", "Die", "Das", "Ein", "Eine", "Les", "Le", "La"}
		sub.Subsonic.Artists.IgnoredArticles = strings.Join(ignoredArticles, " ")

		stripArticle := func(name string) string {
			for _, article := range ignoredArticles {
				prefix := article + " "
				if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
					return strings.TrimSpace(name[len(prefix):])
				}
			}
			return name
		}

		artistMap := make(map[string]types.SubsonicArtist)
		for _, song := range songMap {
			idStr := fmt.Sprintf("%d", song.ArtistID)
			if _, exists := artistMap[idStr]; !exists {
				artistMap[idStr] = types.SubsonicArtist{
					ID:       idStr,
					Name:     song.Name,
					CoverArt: song.ProfilePicture,
				}
			}
		}

		alphaIndex := make(map[string][]types.SubsonicArtist)
		for _, artist := range artistMap {
			stripped := stripArticle(artist.Name)
			runes := []rune(stripped)
			initial := "#"
			if len(runes) > 0 && unicode.IsLetter(runes[0]) {
				initial = strings.ToUpper(string(runes[0]))
			}
			alphaIndex[initial] = append(alphaIndex[initial], artist)
		}

		var keys []string
		for k := range alphaIndex {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			sort.Slice(alphaIndex[key], func(i, j int) bool {
				a := stripArticle(alphaIndex[key][i].Name)
				b := stripArticle(alphaIndex[key][j].Name)
				return strings.ToLower(a) < strings.ToLower(b)
			})

			sub.Subsonic.Artists.Index = append(sub.Subsonic.Artists.Index, types.SubsonicArtistIndexItem{
				Name:   key,
				Artist: alphaIndex[key],
			})
		}

		w.Header().Set("Cache-Control", "public, max-age=3600, immutable")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Artist", "yes")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sub)

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

		if err := json.Unmarshal(body, &playback); err != nil {
			http.Error(w, "failed to parse playback info", config.StatusInternalServerError)
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(playback.Manifest)

		if err != nil {
			http.Error(w, "failed to decode manifest", config.StatusInternalServerError)
			return
		}

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

	// -------------------- Scrobble --------------------
	case rest.Scrobble():
		return

	}
}
