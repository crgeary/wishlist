package utils

func Optional(s string, d string) string {
	if s != "" {
		return s
	}

	return d
}
