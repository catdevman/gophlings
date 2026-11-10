// Package text is named text so it does not clash with the strings standard library.
package text

// Strings are UTF-8 byte sequences. len counts bytes, not characters.
// unicode/utf8.RuneCountInString counts code points.
// https://go.dev/blog/strings
//
// Hello returns the greeting "Hello, 世界".
func Hello() string {
	return ""
}

// ByteLen returns how many bytes s occupies.
func ByteLen(s string) int {
	return 0
}

// RuneCount returns how many Unicode code points s contains.
func RuneCount(s string) int {
	return 0
}
