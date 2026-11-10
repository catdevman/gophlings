package assert

// A type assertion reads the concrete value stored in an interface.
// The comma-ok form reports whether the type matched.
// A type switch compares the dynamic type against several cases.
// https://go.dev/tour/methods/1

// AsString returns the string stored in v and true.
// When v is not a string, return "" and false.
func AsString(v any) (string, bool) {
	return "", false
}

// Number returns the int in v.
// Accept an int, an int64, or a float64 truncated toward zero.
// Return 0 for any other type.
func Number(v any) int {
	return 0
}
