package slice_test

import (
	"reflect"
	"testing"

	"github.com/catdevman/gophlings/17_slice"
)

func TestDefaults(t *testing.T) {
	s := []int{1, 2, 3, 4}
	if got := slice.Left(s); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("Left() = %#v, want [1 2]", got)
	}
	if got := slice.Right(s); !reflect.DeepEqual(got, []int{3, 4}) {
		t.Fatalf("Right() = %#v, want [3 4]", got)
	}
	if got := slice.Middle(s); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Fatalf("Middle() = %#v, want [2 3]", got)
	}
	if got := slice.All(s); !reflect.DeepEqual(got, s) {
		t.Fatalf("All() = %#v, want %#v", got, s)
	}
}

func TestNil(t *testing.T) {
	if got := slice.NilSlice(); got != nil {
		t.Fatalf("NilSlice() = %#v, want a nil slice", got)
	}
}

func TestMake(t *testing.T) {
	got := slice.MakeSlice()
	if len(got) != 3 || cap(got) != 5 {
		t.Fatalf("len=%d cap=%d, want len 3 cap 5", len(got), cap(got))
	}
	if got[0] != 0 || got[1] != 0 || got[2] != 0 {
		t.Fatalf("MakeSlice() = %#v, want three zero ints", got)
	}
}

func TestBoard(t *testing.T) {
	want := [][]string{{"X", "O"}, {"O", "X"}}
	if got := slice.Board(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Board() = %#v, want %#v", got, want)
	}
}

func TestGrow(t *testing.T) {
	got := slice.Grow([]int{1}, 2, 3)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("Grow() = %#v, want [1 2 3]", got)
	}
	got = slice.Grow(nil, 1)
	if !reflect.DeepEqual(got, []int{1}) {
		t.Fatalf("Grow(nil, 1) = %#v, want [1]", got)
	}
}
