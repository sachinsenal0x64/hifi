package middleware

import (
	"encoding/json"
	"hifi/types"
	"net/http"
)

func getArtists(user string, w http.ResponseWriter) {
	sub := types.MetaBanner()
	sub.Subsonic.Artists = &types.SubsonicArtists{}

	artistsMu.RLock()
	ids := append([]int(nil), artistsOrder[user]...)
	userArtists := artistsCache[user]
	artistsMu.RUnlock()

	artists := make([]types.SubsonicArtist, 0, len(ids))
	for _, id := range ids {
		if a, ok := userArtists[id]; ok {
			artists = append(artists, a)
		}
	}

	sub.Subsonic.Artists.Index = []types.SubsonicArtistIndexItem{
		{Artist: artists},
	}

	_ = json.NewEncoder(w).Encode(sub)
}
