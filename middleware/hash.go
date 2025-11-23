package middleware

import (
	"crypto/md5"
	"encoding/hex"
)

// -------------------- HASH --------------------

func Token(passWord, salt string) string {
	data := []byte(passWord + salt)
	hash := md5.Sum(data)
	token := hex.EncodeToString(hash[:])
	return token

}
