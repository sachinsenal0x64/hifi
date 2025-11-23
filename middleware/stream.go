package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/types"
	"io"
	"net/http"
	"net/url"
)

func stream(id string, w http.ResponseWriter, r *http.Request) {

	var playback types.PlaybackInfo
	var manifest types.ManifestData

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
}
