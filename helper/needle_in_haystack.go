package helper

// Digunakan untuk melakukan pencarian value dalam slice
func FindNeedle(haystack []string, needle string) bool {
	for _, element := range haystack {
		if element == needle {
			return true
		}
	}

	return false
}
