package middleware

import (
	"hifi/config"
)

func TidalAuth(clientID string, clientSecret string, refreshToken string) string {

	return config.AccessToken
}
