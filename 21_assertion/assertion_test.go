package assert_test

import (
	"testing"

	"github.com/catdevman/gophlings/21_assertion"
)

func TestAsString(t *testing.T) {
	got, ok := assert.AsString("go")
	if !ok || got != "go" {
		t.Fatalf("AsString(%q) = %q, %v, want %q, true", "go", got, ok, "go")
	}
	got, ok = assert.AsString(3)
	if ok || got != "" {
		t.Fatalf("AsString(3) = %q, %v, want %q, false", got, ok, "")
	}
}

func TestNumber(t *testing.T) {
	if got := assert.Number(42); got != 42 {
		t.Fatalf("Number(42) = %d, want 42", got)
	}
	if got := assert.Number(int64(7)); got != 7 {
		t.Fatalf("Number(int64(7)) = %d, want 7", got)
	}
	if got := assert.Number(3.9); got != 3 {
		t.Fatalf("Number(3.9) = %d, want 3", got)
	}
	if got := assert.Number("no"); got != 0 {
		t.Fatalf("Number(%q) = %d, want 0", "no", got)
	}
}
