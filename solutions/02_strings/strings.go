// Package text is named text so it does not clash with the strings standard library.
package text

import "unicode/utf8"

// Strings are UTF-8 byte sequences. len counts bytes, not characters.
// https://go.dev/blog/strings
//
// Hello returns the greeting "Hello, 世界".
func Hello() string {
	return "Hello, 世界"
}

// ByteLen returns how many bytes s occupies.
func ByteLen(s string) int {
	return len(s)
}

// RuneCount returns how many Unicode code points s contains.
func RuneCount(s string) int {
	return utf8.RuneCountInString(s)
}
