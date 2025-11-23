package middleware

func firstNonEmpty(s1, s2 string) string {
	if s1 != "" {
		return s1
	}
	return s2
}
