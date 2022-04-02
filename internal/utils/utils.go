package utils

// Used to ignore declared and not used annoying message
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
