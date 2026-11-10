package shapes_test

import (
	"testing"

	"github.com/catdevman/gophlings/20_interface"
)

func TestAlias(t *testing.T) {
	var f float64 = 98.6
	if got := shapes.AsFloat(f); got != f {
		t.Fatalf("AsFloat(%v) = %v, want the same float64", f, got)
	}
}

func TestImplicit(t *testing.T) {
	v := shapes.Vertex{X: 3, Y: 4}
	if got := v.Abs(); got != 5 {
		t.Fatalf("Abs() = %v, want 5", got)
	}
	if got := shapes.Measure(v); got != 5 {
		t.Fatalf("Measure() = %v, want 5", got)
	}
}

func TestTypedNil(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Fatal("NilSpeaker panicked: return a nil *Quiet stored in a Speaker, not a nil interface")
		}
	}()
	s := shapes.NilSpeaker()
	if s == nil {
		t.Fatal("NilSpeaker is a nil interface; it should hold a nil *Quiet")
	}
	if got := s.Speak(); got != "..." {
		t.Fatalf("Speak() = %q, want %q", got, "...")
	}
}

func TestNilInterface(t *testing.T) {
	s := shapes.MissingSpeaker()
	if s != nil {
		t.Fatal("MissingSpeaker should be a nil interface")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("calling a method on a nil interface should panic")
		}
	}()
	s.Speak()
}

func TestQuiet(t *testing.T) {
	q := &shapes.Quiet{}
	if got := q.Speak(); got != "hi" {
		t.Fatalf("Speak() = %q, want %q", got, "hi")
	}
}

func TestEmpty(t *testing.T) {
	if got := shapes.Pack(42); got != 42 {
		t.Fatalf("Pack(42) = %#v, want 42", got)
	}
	if got := shapes.Pack("go"); got != "go" {
		t.Fatalf("Pack(%q) = %#v, want %q", "go", got, "go")
	}
}
