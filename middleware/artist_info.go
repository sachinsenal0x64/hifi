package middleware

import (
	"encoding/json"
	"hifi/types"
	"net/http"
)

func getArtistInfo(id string, w http.ResponseWriter) {

	artistInfoMu.RLock()
	info, found := artistInfoCache[id]
	artistInfoMu.RUnlock()

	if !found {

		info = types.SubsonicArtistInfo{
			ID: id,
		}
		artistInfoMu.Lock()
		artistInfoCache[id] = info
		artistInfoMu.Unlock()
	}

	sub := types.MetaBanner()
	sub.Subsonic.ArtistInfo = &info

	json.NewEncoder(w).Encode(sub)
}
