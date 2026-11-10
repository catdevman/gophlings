package numeric_test

import (
	"os"
	"strings"
	"testing"

	"github.com/catdevman/gophlings/10_numeric"
)

func TestShifts(t *testing.T) {
	code := source(t)
	if !strings.Contains(code, "<<") || !strings.Contains(code, ">>") {
		t.Fatal("define Big and Small with << and >>")
	}
}

func TestSmallInt(t *testing.T) {
	if got := numeric.UseSmallInt(); got != 21 {
		t.Fatalf("UseSmallInt() = %d, want 21", got)
	}
}

func TestSmallFloat(t *testing.T) {
	if got := numeric.UseSmallFloat(); got != 0.2 {
		t.Fatalf("UseSmallFloat() = %v, want 0.2", got)
	}
}

func TestBigFloat(t *testing.T) {
	const want = 1.2676506002282295e+29
	if got := numeric.UseBigFloat(); got != want {
		t.Fatalf("UseBigFloat() = %v, want %v", got, want)
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("numeric.go")
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
