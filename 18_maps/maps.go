package dict

// A map stores a value under a key. A missing key yields the zero value,
// so lookup uses the comma-ok form to tell a missing key from a stored zero.
// delete removes a key. https://go.dev/tour/moretypes/1

// Scores returns a map with "ada" set to 10 and "lin" set to 8.
func Scores() map[string]int {
	return nil
}

// Set stores value under key.
func Set(m map[string]int, key string, value int) {
}

// Remove deletes key from m.
func Remove(m map[string]int, key string) {
}

// Lookup reports the value and whether key is present.
func Lookup(m map[string]int, key string) (int, bool) {
	return 0, false
}
