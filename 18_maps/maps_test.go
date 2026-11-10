package dict_test

import (
	"testing"

	"github.com/catdevman/gophlings/18_maps"
)

func TestScores(t *testing.T) {
	m := dict.Scores()
	if m["ada"] != 10 || m["lin"] != 8 {
		t.Fatalf("Scores() = %#v, want ada:10 and lin:8", m)
	}
}

func TestMutate(t *testing.T) {
	m := dict.Scores()
	dict.Set(m, "lin", 9)
	dict.Set(m, "grace", 7)
	dict.Remove(m, "ada")

	if _, ok := dict.Lookup(m, "ada"); ok {
		t.Fatal("ada should have been deleted")
	}
	if got, ok := dict.Lookup(m, "lin"); !ok || got != 9 {
		t.Fatalf("lin = %d, %v, want 9, true", got, ok)
	}
	if got, ok := dict.Lookup(m, "grace"); !ok || got != 7 {
		t.Fatalf("grace = %d, %v, want 7, true", got, ok)
	}
	if got, ok := dict.Lookup(m, "missing"); ok || got != 0 {
		t.Fatalf("missing = %d, %v, want 0, false", got, ok)
	}
}
