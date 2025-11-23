package middleware

import (
	"encoding/json"
	"hifi/types"
	"net/http"
)

func song(id string, w http.ResponseWriter) {

	songMu.RLock()
	song, exists := songMap[id]
	songMu.RUnlock()

	if !exists {
		http.Error(w, "song not found", http.StatusNotFound)
		return
	}

	sub := types.MetaBanner()
	sub.Subsonic.Song = &song

	json.NewEncoder(w).Encode(sub)
}
