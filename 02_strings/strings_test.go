package text_test

import (
	"testing"
	"unicode/utf8"

	"github.com/catdevman/gophlings/02_strings"
)

func TestHello(t *testing.T) {
	const want = "Hello, 世界"
	if got := text.Hello(); got != want {
		t.Fatalf("Hello() = %q, want %q", got, want)
	}
}

func TestUTF8(t *testing.T) {
	s := text.Hello()
	if s == "" {
		t.Fatal("Hello() is empty")
	}
	if got := text.ByteLen(s); got != len(s) {
		t.Fatalf("ByteLen(%q) = %d, want %d bytes", s, got, len(s))
	}
	if got := text.RuneCount(s); got != utf8.RuneCountInString(s) {
		t.Fatalf("RuneCount(%q) = %d, want %d runes", s, got, utf8.RuneCountInString(s))
	}
	if text.ByteLen(s) == text.RuneCount(s) {
		t.Fatal("Hello, 世界 uses more bytes than runes; len counts bytes")
	}
}
