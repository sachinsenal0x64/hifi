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
	"syscall"
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

		if rangeHeader := r.Header.Get("Range"); rangeHeader != "" {
			req.Header.Set("Range", rangeHeader)
			fmt.Println(rangeHeader)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Proxy request failed", http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		for k, vv := range res.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}

		w.Header().Set(config.HeaderContentType, "audio/flac")
		w.Header().Set("Accept-Ranges", "bytes")
		w.WriteHeader(res.StatusCode)

		buf := make([]byte, 64*1024)
		for {
			n, err := res.Body.Read(buf)
			if n > 0 {
				_, writeErr := w.Write(buf[:n])
				if writeErr != nil {
					if errors.Is(writeErr, syscall.EPIPE) || errors.Is(writeErr, syscall.ECONNRESET) {
						return
					}
					fmt.Printf("Write error: %v\n", writeErr)
					return
				}

				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
			}

			if err != nil {
				if err != io.EOF {
					fmt.Printf("Read error: %v\n", err)
				}
				break
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
