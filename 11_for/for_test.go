package loops_test

import (
	"os"
	"strings"
	"testing"

	"github.com/catdevman/gophlings/11_for"
)

func TestSumTo(t *testing.T) {
	if got := loops.SumTo(4); got != 10 {
		t.Fatalf("SumTo(4) = %d, want 10", got)
	}
	if got := loops.SumTo(0); got != 0 {
		t.Fatalf("SumTo(0) = %d, want 0", got)
	}
}

func TestSumWhile(t *testing.T) {
	if got := loops.SumWhile(4); got != 10 {
		t.Fatalf("SumWhile(4) = %d, want 10", got)
	}
}

func TestFind(t *testing.T) {
	if got := loops.Find([]int{5, 6, 7}, 6); got != 1 {
		t.Fatalf("Find(6) = %d, want 1", got)
	}
	if got := loops.Find([]int{5, 6, 7}, 9); got != -1 {
		t.Fatalf("Find(9) = %d, want -1", got)
	}
	if got := loops.Find(nil, 1); got != -1 {
		t.Fatalf("Find(nil, 1) = %d, want -1", got)
	}
}

func TestSumRange(t *testing.T) {
	if got := loops.SumRange([]int{1, 2, 3}); got != 6 {
		t.Fatalf("SumRange() = %d, want 6", got)
	}
}

func TestUsesFor(t *testing.T) {
	code := source(t)
	if strings.Count(code, "for ") < 4 {
		t.Fatal("write a for loop in SumTo, SumWhile, Find, and SumRange")
	}
	if !strings.Contains(code, "range ") && !strings.Contains(code, "range\t") {
		t.Fatal("walk nums with range")
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("for.go")
	if err != nil {
		t.Fatal(err)
	}
	var bld strings.Builder
	for _, line := range strings.Split(string(b), "\n") {
		if i := strings.Index(line, "//"); i >= 0 {
			line = line[:i]
		}
		bld.WriteString(line)
		bld.WriteByte('\n')
	}
	return bld.String()
}
