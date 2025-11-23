package middleware

import (
	"encoding/json"
	"hifi/config"
	"hifi/types"
	"io"
	"log/slog"
	"net/http"
)

func GetNew() []int {

	var tidalNew types.TidalNew

	tidalURL := QueryBuild(config.TidalHost, "/v1/pages/explore_new_music")

	q := tidalURL.Query()
	q.Set("countryCode", "US")
	q.Set("deviceType", "BROWSER")
	tidalURL.RawQuery = q.Encode()

	req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
	req.Header.Set("Authorization", "Bearer "+TidalAuth())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("failed to send request to Tidal", "error", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read Tidal response", "error", err)
		return nil
	}

	if err := json.Unmarshal(body, &tidalNew); err != nil {
		slog.Error("failed to parse Tidal response", "error", err)
		return nil
	}

	return extractIDs(&tidalNew, 2)
}
