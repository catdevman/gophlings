package vertex_test

import (
	"testing"

	"github.com/catdevman/gophlings/15_struct"
)

func TestFields(t *testing.T) {
	v := vertex.New(2, 5)
	if v.X != 2 || v.Y != 5 {
		t.Fatalf("New(2, 5) = %+v, want {X:2 Y:5}", v)
	}
}

func TestPointerField(t *testing.T) {
	v := vertex.New(1, 1)
	vertex.Move(&v, 2, 3)
	if v.X != 3 || v.Y != 4 {
		t.Fatalf("after Move, vertex = %+v, want {X:3 Y:4}", v)
	}
}

func TestValueReceiver(t *testing.T) {
	if got := vertex.New(1, 2).Sum(); got != 3 {
		t.Fatalf("Sum() = %d, want 3", got)
	}
}

func TestPointerReceiver(t *testing.T) {
	v := vertex.New(2, 3)
	v.Scale(4)
	if v.X != 8 || v.Y != 12 {
		t.Fatalf("after Scale, vertex = %+v, want {X:8 Y:12}", v)
	}
}
