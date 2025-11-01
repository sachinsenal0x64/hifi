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
	"sync"
	"time"
	"unicode"
)

// -------------------- REWRITE --------------------

var (
	query    = make(map[string]string)
	songMap  = make(map[string]types.SubsonicSong)
	coverMap = make(map[string]string)
	album    = make(map[string]string)

	playback types.PlaybackInfo
	manifest types.ManifestData

	tidalSearch types.TidalSearchResponse
	tidalArtist types.TidalArtistResponse

	queryMu sync.RWMutex
	songMu  sync.RWMutex
	coverMu sync.RWMutex
	albumMu sync.RWMutex

	userTidalAlbums = make(map[string]types.TidalAlbumResponse)
	userAlbumMu     sync.RWMutex
)

func RewriteRequest(w http.ResponseWriter, r *http.Request) {

	user := r.URL.Query().Get("u")
	search := r.URL.Query().Get("query")

	switch r.URL.Path {
	case rest.Search3View():

		if search != "" {
			queryMu.Lock()
			query[user] = search
			queryMu.Unlock()
		}

		queryMu.RLock()
		qu := query[user]
		queryMu.RUnlock()

		// Tidal search URL
		tidalURL := &url.URL{
			Scheme: config.Scheme,
			Host:   config.TidalHost,
			Path:   "/v1/search/tracks",
		}
		q := tidalURL.Query()
		q.Set("query", qu)
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

		artistMap := make(map[int]bool)
		albumMap := make(map[int]bool)

		for _, item := range tidalSearch.Items {

			albumID := fmt.Sprint(item.Album.ID)
			songID := fmt.Sprint(item.ID)

			coverUUID := item.Album.Cover

			coverMu.Lock()
			coverMap[albumID] = coverUUID
			coverMap[songID] = coverUUID
			coverMu.Unlock()

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
					Year:     item.StreamStartDate[0:4],
					IsDir:    true,
					Duration: item.Duration,
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

			songMu.Lock()
			songMap[songID] = song
			songMu.Unlock()

		}

		// Write response
		json.NewEncoder(w).Encode(sub)

	// -------------------- COVER ART --------------------

	case rest.GetCoverArtView():

		id := r.URL.Query().Get("id")
		size := r.URL.Query().Get("size")

		sizeMapping := map[int]int{
			100: 160,
			200: 320,
			300: 320, // 750x750
			450: 640, // 1080x1080
			500: 750, // 1280x1280
		}

		s, _ := strconv.ParseInt(size, 10, 64)
		mappedSize := sizeMapping[int(s)]

		redirectURL := fmt.Sprintf(
			"%s://%s/images/%s/%dx%d.jpg",
			config.Scheme,
			config.TidalStaticHost,
			FormatCoverID(id),
			mappedSize,
			mappedSize,
		)

		http.Redirect(w, r, redirectURL, config.StatusRedirectPermanent)

	// -------------------- getSong --------------------

	case rest.GetSong():
		id := r.URL.Query().Get("id")

		songMu.RLock()
		song := songMap[id]
		songMu.RUnlock()

		sub := types.MetaBanner()
		sub.Subsonic.Song = &song

		json.NewEncoder(w).Encode(sub)

	// -------------------- getAlbumList2 --------------------
	case rest.GetAlbumList2View():
		albumMu.RLock()
		ids := album[user]
		albumMu.RUnlock()

		sub := types.MetaBanner()
		var subsonicAlbum types.SubsonicAlbum
		subsonicAlbum.ID = ids
		subsonicAlbum.IsDir = true
		subsonicAlbum.Created = time.Now().Format(time.RFC3339)

		userAlbumData := userTidalAlbums[user]

		for _, item := range userAlbumData.Items {
			song := types.SubsonicSong{
				ID:       fmt.Sprint(item.Item.ID),
				Duration: item.Item.Duration,
				Title:    item.Item.Title,
				Album:    item.Item.Album.Title,
				Artist:   item.Item.Artist.Name,
				CoverArt: item.Item.Album.Cover,
				Parent:   item.Item.Artist.ID,
				Type:     "music",
				IsVideo:  false,
				Suffix:   "flac",
				Year:     item.Item.StreamStartDate[0:4],
				Track:    item.Item.TrackNumber,
				ArtistID: fmt.Sprint(item.Item.Artist.ID),
				AlbumID:  fmt.Sprint(item.Item.Album.ID),
			}

			subsonicAlbum.Parent = item.Item.Artist.ID
			subsonicAlbum.ArtistID = item.Item.Artist.ID
			subsonicAlbum.Name = item.Item.Album.Title
			subsonicAlbum.Title = item.Item.Album.Title
			subsonicAlbum.Artist = item.Item.Artist.Name
			subsonicAlbum.CoverArt = item.Item.Album.Cover
			subsonicAlbum.Year = item.Item.StreamStartDate[0:4]
			subsonicAlbum.SongCount = len(userAlbumData.Items)
			subsonicAlbum.Duration += item.Item.Duration
			subsonicAlbum.Song = append(subsonicAlbum.Song, song)

			songMu.Lock()
			songMap[song.ID] = song
			songMu.Unlock()
		}

		sub.Subsonic.AlbumList2 = &types.SubsonicAlbumList{
			Album: []types.SubsonicAlbum{subsonicAlbum},
		}

		json.NewEncoder(w).Encode(sub)

	// -------------------- getAlbum --------------------
	case rest.GetAlbumView():
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing album id", http.StatusBadRequest)
			return
		}

		albumMu.Lock()
		album[user] = id
		albumMu.Unlock()

		fmt.Printf("[Get Album Request] User: %s | Album ID: %s\n", user, id)

		// Build Tidal request
		tidalURL := &url.URL{
			Scheme: config.Scheme,
			Host:   config.TidalHost,
			Path:   fmt.Sprintf("/v1/albums/%s/items", id),
		}
		q := tidalURL.Query()
		q.Set("countryCode", "US")
		q.Set("limit", "100")
		q.Set("offset", "0")
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

		// Local struct — don't use any global tidalAlbum
		var tidalAlbum types.TidalAlbumResponse
		if err := json.Unmarshal(body, &tidalAlbum); err != nil {
			http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
			return
		}

		// ✅ Store album data per user to avoid shared state
		userAlbumMu.Lock()
		userTidalAlbums[user] = tidalAlbum
		userAlbumMu.Unlock()

		// Prepare Subsonic response
		sub := types.MetaBanner()
		albumResp := &types.SubsonicAlbum{}

		for _, item := range tidalAlbum.Items {
			song := types.SubsonicSong{
				ID:       fmt.Sprint(item.Item.ID),
				Duration: item.Item.Duration,
				Title:    item.Item.Title,
				Album:    item.Item.Album.Title,
				Artist:   item.Item.Artist.Name,
				CoverArt: item.Item.Album.Cover,
				Parent:   item.Item.Artist.ID,
				Type:     "music",
				IsVideo:  false,
				Suffix:   "flac",
				Year:     item.Item.StreamStartDate[0:4],
				Track:    item.Item.TrackNumber,
				ArtistID: fmt.Sprint(item.Item.Artist.ID),
				AlbumID:  fmt.Sprint(item.Item.Album.ID),
			}

			albumResp.ID = id
			albumResp.Parent = item.Item.Artist.ID
			albumResp.ArtistID = item.Item.Artist.ID
			albumResp.IsDir = true
			albumResp.Name = item.Item.Album.Title
			albumResp.Title = item.Item.Album.Title
			albumResp.Artist = item.Item.Artist.Name
			albumResp.CoverArt = item.Item.Album.Cover
			albumResp.Year = item.Item.StreamStartDate[0:4]
			albumResp.SongCount = len(tidalAlbum.Items)
			albumResp.Duration += item.Item.Duration
			albumResp.Song = append(albumResp.Song, song)

			songMu.Lock()
			songMap[song.ID] = song
			songMu.Unlock()
		}

		sub.Subsonic.Album = albumResp
		json.NewEncoder(w).Encode(sub)

	// -------------------- getArtists --------------------

	case rest.GetArtistsView():

		queryMu.RLock()
		qu := query[user]
		queryMu.RUnlock()

		fmt.Println(qu)

		// Tidal search URL
		tidalURL := &url.URL{
			Scheme: config.Scheme,
			Host:   config.TidalHost,
			Path:   "/v1/search/artists",
		}
		q := tidalURL.Query()
		q.Set("query", qu)
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

		if err := json.Unmarshal(body, &tidalArtist); err != nil {
			http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
			return
		}

		sub := types.MetaBanner()
		sub.Subsonic.Artists = &types.SubsonicArtists{}

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
		for _, song := range tidalArtist.Items {
			idStr := fmt.Sprintf("%d", song.ID)
			if _, exists := artistMap[idStr]; !exists {
				artistMap[idStr] = types.SubsonicArtist{
					ID:       idStr,
					Name:     song.Name,
					CoverArt: song.Picture,
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
		q.Set("audioquality", "HI_RES_LOSSLESS")
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
