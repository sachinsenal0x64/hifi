package middleware

import (
	"fmt"
	"hifi/config"
	"hifi/routes/rest"
	"io"
	"net/http"
	"net/url"
	"os"
)

// -------------------- REWRITE --------------------

func RewriteRequest(r *http.Request) {
	r.URL.Scheme = config.Scheme
	r.URL.Host = config.TidalHost
	r.Host = config.TidalHost

	// Add Tidal auth header
	tidalToken := TidalAuth()

	switch r.URL.Path {
	case rest.Search3View():
		query := r.URL.Query().Get("query")
		if query != "" {
			tidalURL := &url.URL{
				Scheme: config.Scheme,
				Host:   config.TidalHost,
				Path:   "/v1/search/tracks",
			}

			q := tidalURL.Query()
			q.Set("query", query)
			q.Set("limit", "25")
			q.Set("offset", "0")
			q.Set("countryCode", "US")
			tidalURL.RawQuery = q.Encode()

			r.URL = tidalURL
			r.Host = tidalURL.Host

			fmt.Println(tidalURL)

			req, err := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
			req.Header.Set("Authorization", "Bearer "+tidalToken)

			if err != nil {
				fmt.Printf("client: could not create request: %s\n", err)
				os.Exit(1)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("client: error making http request: %s\n", err)
				os.Exit(1)
			}

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Printf("client: could not read response body: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("client: response body: %s\n", resBody)

		}
	}
}
