package config

import (
	"os"
)

// Subsonic User
type Person struct {
	UserName string
	PassWord string
}

var (

	// Blacklist
	ExcludedPaths = []string{
		"/admin/home",
	}

	// ENV
	AccessToken = os.Getenv("ACCESS_TOKEN")

	// Target Host
	TargetHost = os.Getenv("TARGET_HOST")
)

const (

	// TIDAL API
	Scheme          = "https"
	TidalHost       = "api.tidal.com"
	TidalStaticHost = "resources.tidal.com"

	// Hifi Server
	Host = "127.0.0.1"
	Port = "5000"

	// HTTP methods
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodDelete  = "DELETE"
	MethodPatch   = "PATCH"
	MethodOptions = "OPTIONS"

	// Content types
	ContentTypeJSON = "application/json"

	// Common headers
	HeaderContentType  = "Content-Type"
	HeaderCacheControl = "Cache-Control"
	HeaderConnection   = "Connection"
	HeaderAllowOrigin  = "Access-Control-Allow-Origin"
	HeaderAllowMethods = "Access-Control-Allow-Methods"
	HeaderAllowHeaders = "Access-Control-Allow-Headers"
	Authorization      = "Authorization"

	// CORS
	CORSAllowOrigin = "*"

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
