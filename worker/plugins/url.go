package plugins

import (
	"net/url"
	"worker/config"
)

func URLBuild(host, path string) string {
	return (&url.URL{
		Scheme: config.Scheme,
		Host:   host,
		Path:   path,
	}).String()
}
