package conditional_test

import (
	"os"
	"strings"
	"testing"

	"github.com/catdevman/gophlings/12_conditional"
)

func TestSign(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{2, "positive"},
		{-3, "negative"},
		{0, "zero"},
	}
	for _, tc := range cases {
		if got := conditional.Sign(tc.n); got != tc.want {
			t.Fatalf("Sign(%d) = %q, want %q", tc.n, got, tc.want)
		}
	}
}

func TestAbs(t *testing.T) {
	if got := conditional.Abs(-4); got != 4 {
		t.Fatalf("Abs(-4) = %d, want 4", got)
	}
	if got := conditional.Abs(4); got != 4 {
		t.Fatalf("Abs(4) = %d, want 4", got)
	}
}

func TestHalfIfEven(t *testing.T) {
	if got := conditional.HalfIfEven(8); got != 4 {
		t.Fatalf("HalfIfEven(8) = %d, want 4", got)
	}
	if got := conditional.HalfIfEven(7); got != 7 {
		t.Fatalf("HalfIfEven(7) = %d, want 7", got)
	}
}

func TestDay(t *testing.T) {
	if got := conditional.Day(1); got != "mon" {
		t.Fatalf("Day(1) = %q, want mon", got)
	}
	if got := conditional.Day(2); got != "tue" {
		t.Fatalf("Day(2) = %q, want tue", got)
	}
	if got := conditional.Day(9); got != "other" {
		t.Fatalf("Day(9) = %q, want other", got)
	}
}

func TestSyntax(t *testing.T) {
	code := source(t)
	if !strings.Contains(code, "else") {
		t.Fatal("use else")
	}
	if !strings.Contains(code, "switch ") && !strings.Contains(code, "switch\t") {
		t.Fatal("use switch in Day")
	}
	if !strings.Contains(code, ":=") {
		t.Fatal("use a short statement in the if, such as if half := n / 2; ...")
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("conditional.go")
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
