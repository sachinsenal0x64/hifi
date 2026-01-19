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
	var host string

	mode := config.MODE

	if mode == "managed" {
		host = config.ManageHost

	} else {
		host = config.TidalHost
	}

	tidalURL := &url.URL{
		Scheme: config.Scheme,
		Host:   host,
		Path:   fmt.Sprintf("/v1/tracks/%s/playbackinfopostpaywall/v4", id),
	}

	q := tidalURL.Query()
	q.Set("audioquality", "LOSSLESS")
	q.Set("playbackmode", "STREAM")
	q.Set("assetpresentation", "FULL")
	tidalURL.RawQuery = q.Encode()

	if config.MODE == "managed" {

		req, err := http.NewRequest(http.MethodGet, tidalURL.String(), nil)
		if err != nil {
			http.Error(w, "Proxy request creation failed", http.StatusInternalServerError)
			return
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Proxy request failed", http.StatusInternalServerError)
			return
		}

		defer res.Body.Close()

		w.WriteHeader(res.StatusCode)

		_, err = io.Copy(w, res.Body)
		if err != nil {
			fmt.Printf("Error streaming to client: %v\n", err)
		}
		return

	} else {
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

}
