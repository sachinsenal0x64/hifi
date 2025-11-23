package middleware

import (
	"regexp"
	"strings"
)

func cleanHTML(input string) string {
	input = strings.ReplaceAll(input, "<br/>", "\n")

	re := regexp.MustCompile(`<[^>]*>`)
	cleaned := re.ReplaceAllString(input, "")

	return cleaned
}
