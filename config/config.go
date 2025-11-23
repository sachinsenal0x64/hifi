package config

import (
	"hifi/routes/rest"
	"os"
)

var (

	// Whitelist
	ValidPaths = []string{
		rest.Ping(),
		rest.Search3View(),
		rest.GetArtistsView(),
		rest.GetCoverArtView(),
		rest.Stream(),
		rest.GetSong(),
		rest.Scrobble(),
		rest.GetAlbumView(),
		rest.GetAlbumList2View(),
		rest.GetArtistInfoView(),
		rest.GetArtistView(),
		rest.Fresh(),
		rest.GetTopSongs(),
	}

	// ENV

	ClientID     = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")

	// TIDAL API
	TidalHost    = os.Getenv("TIDAL_HOST")
	RefreshToken = os.Getenv("TIDAL_REFRESH")

	// CORS
	CORSAllowOrigin = "*"

	// Server port and fallback port
	Port = []string{"8080", "5000", "5005"}

	// Hifi Server
	Host = "0.0.0.0"

	// Hifi Scheme
	HifiScheme = "http"
)

const (

	// TIDAL API
	Scheme          = "https"
	TidalStaticHost = "resources.tidal.com"
	TidalAuthHost   = "auth.tidal.com"

	// HTTP methods
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodDelete  = "DELETE"
	MethodPatch   = "PATCH"
	MethodOptions = "OPTIONS"

	// Content types
	ContentTypeJSON = "application/json"
	ContentTypeForm = "application/x-www-form-urlencoded"

	// Common headers
	HeaderContentType  = "Content-Type"
	HeaderCacheControl = "Cache-Control"
	HeaderConnection   = "Connection"
	HeaderAllowOrigin  = "Access-Control-Allow-Origin"
	HeaderAllowMethods = "Access-Control-Allow-Methods"
	HeaderAllowHeaders = "Access-Control-Allow-Headers"
	Authorization      = "Authorization"

	// HTTP State codes
	StatusOK                  = 200
	StatusCreated             = 201
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusNotFound            = 404
	StatusInternalServerError = 500
	StatusMethodNotAllowed    = 405
	StatusMultipleChoices     = 300
	StatusRedirectPermanent   = 308
)
