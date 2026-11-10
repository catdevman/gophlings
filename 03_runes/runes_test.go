package runes_test

import (
	"testing"

	"github.com/catdevman/gophlings/03_runes"
)

func TestLetterG(t *testing.T) {
	if got := runes.LetterG(); got != 'G' {
		t.Fatalf("LetterG() = %q, want 'G'", got)
	}
}

func TestDelta(t *testing.T) {
	if got := runes.Delta(); got != 'δ' {
		t.Fatalf("Delta() = %q, want 'δ'", got)
	}
}

func TestFirst(t *testing.T) {
	if got := runes.First("世界"); got != '世' {
		t.Fatalf("First(%q) = %q, want '世'", "世界", got)
	}
	if got := runes.First(""); got != 0 {
		t.Fatalf("First(\"\") = %q, want 0", got)
	}
}
