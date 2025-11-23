package middleware

import (
	"hifi/config"
	"net/url"
)

func QueryBuild(host, path string) *url.URL {
	return &url.URL{
		Scheme: config.Scheme,
		Host:   host,
		Path:   path,
	}
}
