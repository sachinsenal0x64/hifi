package middleware

import "encoding/base64"

// -------------------- SALT --------------------

func Salt(length int) string {
	data := make([]byte, length)
	return base64.StdEncoding.EncodeToString(data)
}
