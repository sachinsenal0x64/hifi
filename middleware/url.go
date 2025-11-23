package middleware

import (
	"hifi/config"
	"net/url"
)

func URLBuild(host, path string) string {
	return (&url.URL{
		Scheme: config.Scheme,
		Host:   host,
		Path:   path,
	}).String()
}
