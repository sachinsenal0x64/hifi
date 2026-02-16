package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hifi/config"
	"hifi/types"
	"io"
	"net/http"
	"net/url"
	"strings"
	"syscall"
	"time"
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

	var proxyClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	if config.MODE == "managed" {
		req, err := http.NewRequest(http.MethodGet, tidalURL.String(), nil)
		if err != nil {
			http.Error(w, "Request failed", 500)
			return
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36")

		if rh := r.Header.Get("Range"); rh != "" {
			req.Header.Set("Range", rh)
		}

		res, err := proxyClient.Do(req)
		if err != nil {
			http.Error(w, "Proxy failed", 500)
			return
		}
		defer res.Body.Close()

		for k, vv := range res.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}

		w.Header().Set("Content-Type", "audio/flac")
		w.Header().Set("Accept-Ranges", "bytes")
		w.WriteHeader(res.StatusCode)

		buf := make([]byte, 256*1024)
		_, err = io.CopyBuffer(w, res.Body, buf)

		if err != nil {
			if !errors.Is(err, syscall.EPIPE) && !strings.Contains(err.Error(), "wsasend") {
				fmt.Printf("Stream error: %v\n", err)
			}
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
