package functions

import "testing"

func TestBar(t *testing.T) {
	bar("Important parameter")
}

func TestBaz(t *testing.T) {
	n, err := baz()
	if err != nil {
		t.Fatal(err)
	}
	if n != 42 {
		t.Fatalf("baz() = %d, want 42", n)
	}
}
