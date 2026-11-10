package arrays_test

import (
	"testing"

	"github.com/catdevman/gophlings/16_array"
)

func TestEvens(t *testing.T) {
	got := arrays.Evens()
	want := [3]int{2, 4, 6}
	if got != want {
		t.Fatalf("Evens() = %v, want %v", got, want)
	}
}

func TestLast(t *testing.T) {
	got := arrays.Last([4]string{"a", "b", "c", "d"})
	if got != "d" {
		t.Fatalf("Last() = %q, want %q", got, "d")
	}
}

func TestSame(t *testing.T) {
	if !arrays.Same([2]int{1, 2}, [2]int{1, 2}) {
		t.Fatal("equal arrays should compare equal")
	}
	if arrays.Same([2]int{1, 2}, [2]int{2, 1}) {
		t.Fatal("different arrays should not compare equal")
	}
}
