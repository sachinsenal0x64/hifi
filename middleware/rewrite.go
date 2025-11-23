package middleware

import (
	"hifi/routes/rest"
	"hifi/types"
	"net/http"
	"sync"
)

// -------------------- REWRITE --------------------
var (
	query          = make(map[string]string)
	userAlbumCache = make(map[string]map[string]types.SubsonicAlbum)

	Public = "global"

	queryMu     sync.RWMutex
	albumMu     sync.RWMutex
	useralbumMu sync.RWMutex
)

// -------------------- REWRITE --------------------

func RewriteRequest(w http.ResponseWriter, r *http.Request) {

	var (
		user   = r.URL.Query().Get("u")
		search = r.URL.Query().Get("query")
		id     = r.URL.Query().Get("id")
		size   = r.URL.Query().Get("size")
	)
	switch r.URL.Path {

	// -------------------- Search3 --------------------
	case rest.Search3View():
		search3(search, user, w)

	// -------------------- Cover --------------------
	case rest.GetCoverArtView():
		cover(id, size, w, r)

	// -------------------- getSong --------------------
	case rest.GetSong():
		song(id, w)

	// -------------------- getAlbumList2 --------------------
	case rest.GetAlbumList2View():
		getAlbumList2(user, w)

	// -------------------- getAlbum --------------------
	case rest.GetAlbumView():
		getAlbum(id, user, w)

	// -------------------- getArtist --------------------
	case rest.GetArtistView():
		getArtist(id, user, w)

	// -------------------- getArtists --------------------
	case rest.GetArtistsView():
		getArtists(user, w)

	// -------------------- Stream --------------------
	case rest.Stream():
		stream(id, w, r)

	// -------------------- getArtistInfo --------------------
	case rest.GetArtistInfoView():
		getArtistInfo(id, w)

	// -------------------- Scrobble --------------------
	case rest.Scrobble():
		return
	}

}
