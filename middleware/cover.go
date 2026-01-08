package middleware

import (
	"fmt"
	"hifi/config"
	"net/http"
	"strconv"
)

func cover(id string, size string, w http.ResponseWriter, r *http.Request) {
	sizeMapping := map[int]int{
		0:    750,
		20:   80,
		80:   80,
		100:  160,
		200:  320,
		256:  320,
		300:  320, // 750x750
		450:  640, // 1080x1080
		500:  750, // 1280x1280
		512:  750,
		600:  750,
		2137: 1280,
	}

	s, _ := strconv.ParseInt(size, 10, 64)
	mappedSize := sizeMapping[int(s)]

	redirectURL := fmt.Sprintf(
		"%s://%s/images/%s/%dx%d.jpg",
		config.Scheme,
		config.TidalStaticHost,
		FormatCoverID(id),
		mappedSize,
		mappedSize,
	)

	http.Redirect(w, r, redirectURL, config.StatusRedirectPermanent)

}
