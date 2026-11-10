package zero_test

import (
	"os"
	"strings"
	"testing"

	"github.com/catdevman/gophlings/06_zero"
)

func TestZeros(t *testing.T) {
	i, f, s, b, items := zero.Zeros()
	if i != 0 || f != 0 || s != "" || b != false || items != nil {
		t.Fatalf("Zeros() = (%d, %v, %q, %v, %#v), want (0, 0, \"\", false, nil)", i, f, s, b, items)
	}
	if !strings.Contains(source(t), "var ") {
		t.Fatal("declare the variables with var and no initializer")
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("zero.go")
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
