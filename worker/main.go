package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"worker/types"

	"github.com/syumai/workers"
	"github.com/syumai/workers/cloudflare"
	"github.com/syumai/workers/cloudflare/fetch"
)

var playback types.PlaybackInfo
var manifest types.ManifestData

func TidalAuth() string {
	if token := cloudflare.Getenv("TIDAL_TOKEN"); token != "" {
		return strings.TrimSpace(token)
	}
	return ""
}

func handler(w http.ResponseWriter, r *http.Request) {
	trackID := strings.TrimPrefix(r.URL.Path, "/v1/tracks/")
	trackID = strings.SplitN(trackID, "/", 2)[0]

	if trackID == "" {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	tidalURL := fmt.Sprintf(
		"https://api.tidal.com/v1/tracks/%s/playbackinfopostpaywall/v4?audioquality=LOSSLESS&playbackmode=STREAM&assetpresentation=FULL",
		trackID,
	)

	cli := fetch.NewClient()
	req, err := fetch.NewRequest(r.Context(), http.MethodGet, tidalURL, nil)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+TidalAuth())
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	res, err := cli.Do(req, nil)
	if err != nil {
		http.Error(w, "failed to fetch Tidal API", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "failed to read response", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &playback); err != nil {
		http.Error(w, "failed to parse playback info", http.StatusInternalServerError)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(playback.Manifest)
	if err != nil {
		http.Error(w, "failed to decode manifest", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(decoded, &manifest); err != nil {
		http.Error(w, "failed to parse manifest JSON", http.StatusInternalServerError)
		return
	}

	if len(manifest.Urls) == 0 {
		http.Error(w, "no stream URLs found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, manifest.Urls[0], http.StatusFound)
}

func main() {
	workers.Serve(http.HandlerFunc(handler))
}
