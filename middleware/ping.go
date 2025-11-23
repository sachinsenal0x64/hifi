package middleware

import (
	"encoding/json"
	"hifi/config"
	"hifi/types"
	"net/http"
)

func writeSubsonic(w http.ResponseWriter, status string, code int) {
	var resp types.SubsonicWrapper
	resp.Subsonic.Status = status
	resp.Subsonic.Version = "2.0.0"
	resp.Subsonic.Type = "hifi"
	resp.Subsonic.ServerVersion = "2.0.0"
	resp.Subsonic.OpenSubsonic = true

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	_, _ = w.Write(b)
}
