package middleware

import (
	"encoding/json"
	"hifi/config"
	"hifi/types"
	"maps"
	"net/http"
)

func writeSubsonicv2(w http.ResponseWriter, status string, code int, data map[string]any) {
	var base types.SubsonicWrapper
	base.Subsonic.Status = status
	base.Subsonic.Version = "2.0.0"
	base.Subsonic.Type = "hifi"
	base.Subsonic.ServerVersion = "2.0.0"
	base.Subsonic.OpenSubsonic = true

	b, err := json.Marshal(base)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp map[string]any
	if err := json.Unmarshal(b, &resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sub, ok := resp["subsonic-response"].(map[string]any)
	if !ok {
		http.Error(w, "invalid subsonic wrapper", http.StatusInternalServerError)
		return
	}

	if data != nil {
		maps.Copy(sub, data)
	}

	out, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	w.WriteHeader(code)
	_, _ = w.Write(out)
}
