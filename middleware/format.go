package middleware

import "strings"

func FormatCoverID(uuid string) string {
	return strings.ReplaceAll(uuid, "-", "/")
}
