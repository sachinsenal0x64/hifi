package middleware

import (
	"hifi/config"
)

func TidalAuth() string {

	return config.AccessToken
}
