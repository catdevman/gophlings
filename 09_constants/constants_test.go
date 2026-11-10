package constant_test

import (
	"os"
	"strings"
	"testing"

	"github.com/catdevman/gophlings/09_constants"
)

func TestConstants(t *testing.T) {
	if got := constant.App(); got != "gophlings" {
		t.Fatalf("App() = %q, want %q", got, "gophlings")
	}
	if got := constant.Meaning(); got != 42 {
		t.Fatalf("Meaning() = %d, want 42", got)
	}
	code := source(t)
	if !strings.Contains(code, "const ") && !strings.Contains(code, "const(") {
		t.Fatal("declare Title and Answer with const")
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("constants.go")
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
