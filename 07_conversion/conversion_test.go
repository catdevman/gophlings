package conversion_test

import (
	"reflect"
	"testing"

	"github.com/catdevman/gophlings/07_conversion"
)

func TestTruncate(t *testing.T) {
	if got := conversion.Truncate(3.9); got != 3 {
		t.Fatalf("Truncate(3.9) = %d, want 3", got)
	}
	if got := conversion.Truncate(-2.8); got != -2 {
		t.Fatalf("Truncate(-2.8) = %d, want -2", got)
	}
}

func TestFromInt(t *testing.T) {
	if got := conversion.FromInt(3); got != 3 {
		t.Fatalf("FromInt(3) = %v, want 3", got)
	}
}

func TestAdd(t *testing.T) {
	if got := conversion.Add(2, 0.5); got != 2.5 {
		t.Fatalf("Add(2, 0.5) = %v, want 2.5", got)
	}
}

func TestBytesAndString(t *testing.T) {
	if got := conversion.BytesToString([]byte("go")); got != "go" {
		t.Fatalf("BytesToString = %q, want %q", got, "go")
	}
	got := conversion.StringToBytes("go")
	if !reflect.DeepEqual(got, []byte("go")) {
		t.Fatalf("StringToBytes = %v, want %v", got, []byte("go"))
	}
}
