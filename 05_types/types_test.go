package types_test

import (
	"testing"

	"github.com/catdevman/gophlings/05_types"
)

func TestBasics(t *testing.T) {
	ok, name, count, natural, pi, initial := types.Basics()
	if ok != true {
		t.Fatalf("bool = %v, want true", ok)
	}
	if name != "go" {
		t.Fatalf("string = %q, want %q", name, "go")
	}
	if count != -3 {
		t.Fatalf("int = %d, want -3", count)
	}
	if natural != 3 {
		t.Fatalf("uint = %d, want 3", natural)
	}
	if pi != 3.14 {
		t.Fatalf("float64 = %v, want 3.14", pi)
	}
	if initial != 'A' {
		t.Fatalf("byte = %q, want 'A'", initial)
	}
}
