package conversion

// A conversion changes a value's type. The syntax is T(v).
// Unlike some languages, Go does not convert numeric types for you.
// https://go.dev/tour/basics/1

// Truncate converts f to an int, dropping the fraction toward zero.
func Truncate(f float64) int {
	return int(f)
}

// FromInt converts n to a float64.
func FromInt(n int) float64 {
	return float64(n)
}

// Add converts n and adds it to f.
func Add(n int, f float64) float64 {
	return float64(n) + f
}

// BytesToString converts raw bytes into a string.
func BytesToString(b []byte) string {
	return string(b)
}

// StringToBytes converts s into its underlying bytes.
func StringToBytes(s string) []byte {
	return []byte(s)
}
