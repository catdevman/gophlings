package pointers_test

import (
	"testing"

	"github.com/catdevman/gophlings/14_pointers"
)

func TestSetAndGet(t *testing.T) {
	n := 1
	pointers.Set(&n, 5)
	if n != 5 {
		t.Fatalf("after Set, n = %d, want 5", n)
	}
	if got := pointers.Get(&n); got != 5 {
		t.Fatalf("Get() = %d, want 5", got)
	}
}

func TestNewInt(t *testing.T) {
	p := pointers.NewInt(9)
	if p == nil {
		t.Fatal("NewInt returned nil")
	}
	if *p != 9 {
		t.Fatalf("*NewInt(9) = %d, want 9", *p)
	}
}
