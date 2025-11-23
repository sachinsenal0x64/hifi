package config

import "os"

var (

	// Whitelist
	ValidPaths = []string{
		"/v1/signup",
		"/v1/signin",
	}

	// CORS
	CORSAllowOrigin = "*"

	// Hifi Web Server
	Host = "0.0.0.0"
	Port = []string{"5002", "5006"}

	// Hifi Scheme
	HifiScheme = os.Getenv("HIFI_SCHEME")
	HifiProxy  = os.Getenv("HIFI_PROXY")

	// ENV
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	Scheme    = os.Getenv("SCHEME")
	ProxyHost = os.Getenv("SYNC_HOST")
	ProxyKey  = os.Getenv("SYNC_KEY")
	AppID     = os.Getenv("APP_ID")
	ENV       = os.Getenv("ENV")
)

const (

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

	// Cookie
	Cookies = "Cookie"

	// Common headers
	HeaderContentType           = "Content-Type"
	HeaderCacheControl          = "Cache-Control"
	HeaderConnection            = "Connection"
	HeaderAllowOrigin           = "Access-Control-Allow-Origin"
	HeaderAllowMethods          = "Access-Control-Allow-Methods"
	HeaderAllowHeaders          = "Access-Control-Allow-Headers"
	HeaderContentSecurityPolicy = "Content-Security-Policy"

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
)
