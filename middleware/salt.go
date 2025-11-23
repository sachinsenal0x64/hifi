package middleware

import (
	"encoding/base64"
)

// -------------------- SALT --------------------

func Salt(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}
