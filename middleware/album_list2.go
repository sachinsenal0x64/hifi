package middleware

import (
	"encoding/json"
	"hifi/types"
	"maps"
	"net/http"
)

func getAlbumList2(user string, w http.ResponseWriter) {

	albumsMap := make(map[string]types.SubsonicAlbum)

	useralbumMu.Lock()
	maps.Copy(albumsMap, userAlbumCache["global"])
	maps.Copy(albumsMap, userAlbumCache[user])
	useralbumMu.Unlock()

	if len(albumsMap) == 0 {
		http.Error(w, "no cached albums for this user", http.StatusNotFound)
		return
	}

	// Collect all user albums into a slice
	allAlbums := make([]types.SubsonicAlbum, 0, len(albumsMap))
	for _, alb := range albumsMap {
		allAlbums = append(allAlbums, alb)
	}

	sub := types.MetaBanner()
	sub.Subsonic.AlbumList2 = &types.SubsonicAlbumList{
		Album: allAlbums,
	}

	json.NewEncoder(w).Encode(sub)

}
