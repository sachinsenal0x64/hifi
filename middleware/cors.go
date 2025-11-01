package middleware

import (
	"hifi/config"
	"hifi/routes/rest"
	"net/http"
)

// -------------------- CORS --------------------

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case rest.Search3View(), rest.GetArtistsView(), rest.GetSong(), rest.GetAlbumList2View(), rest.Scrobble(), rest.Stream(), rest.GetCoverArtView(), rest.GetAlbumView():
			w.Header().Set(config.HeaderAllowOrigin, config.CORSAllowOrigin)
		}
		next.ServeHTTP(w, r)
	})
}
