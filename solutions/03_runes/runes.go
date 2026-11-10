package runes

// A rune is one Unicode code point. The type is an alias for int32.
// A rune literal uses single quotes: 'A'.
// https://go.dev/blog/strings

// LetterG returns the single character G.
func LetterG() rune {
	return 'G'
}

// Delta returns the Greek small letter delta, δ.
func Delta() rune {
	return 'δ'
}

// First returns the first code point of s.
// Ranging over a string yields runes, not bytes.
func First(s string) rune {
	for _, r := range s {
		return r
	}
	return 0
}
