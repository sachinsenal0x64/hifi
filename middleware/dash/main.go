package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// structure for Tidal JSON response
type TidalPlaybackInfo struct {
	ManifestMimeType string `json:"manifestMimeType"`
	Manifest         string `json:"manifest"`
}

func dashHandler(w http.ResponseWriter, r *http.Request) {
	// --- query params ---

	tidalHost := os.Getenv("TIDAL_HOST")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing ?id parameter", http.StatusBadRequest)
		return
	}

	quality := r.URL.Query().Get("quality")
	if quality == "" {
		quality = "HI_RES_LOSSLESS"
	}

	// --- Access Token ---
	tidalToken := os.Getenv("ACCESS_TOKEN")

	// --- build Tidal API URL ---
	apiURL := fmt.Sprintf(
		"https://%s/v1/tracks/%s/playbackinfo?audioquality=%s&playbackmode=STREAM&assetpresentation=FULL",
		tidalHost, id, quality,
	)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		http.Error(w, "cannot create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+tidalToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "error contacting Tidal", http.StatusBadGateway)
		log.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "error reading response", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Tidal API error: %d %s\n", resp.StatusCode, string(body))
		http.Error(w, "Tidal API error", http.StatusBadGateway)
		return
	}

	// --- parse JSON response ---
	var info TidalPlaybackInfo
	if err := json.Unmarshal(body, &info); err != nil {
		http.Error(w, "invalid JSON from Tidal", http.StatusInternalServerError)
		return
	}

	// --- decode base64 manifest ---
	data, err := base64.StdEncoding.DecodeString(info.Manifest)
	if err != nil {
		http.Error(w, "error decoding manifest", http.StatusInternalServerError)
		return
	}

	// --- return MPD manifest ---
	w.Header().Set("Content-Type", info.ManifestMimeType)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {
	http.HandleFunc("/dash", dashHandler)

	fmt.Println("DASH server running at http://localhost:8080/dash?id=105504512")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
