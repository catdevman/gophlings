package variables_test

import (
	"os"
	"strings"
	"testing"

	"github.com/catdevman/gophlings/04_variables"
)

func TestInitialized(t *testing.T) {
	if got := variables.Initialized(); got != 7 {
		t.Fatalf("Initialized() = %d, want 7", got)
	}
	if !strings.Contains(source(t), "var ") {
		t.Fatal("declare the initialized variable with var")
	}
}

func TestShort(t *testing.T) {
	if got := variables.Short(); got != "gopher" {
		t.Fatalf("Short() = %q, want %q", got, "gopher")
	}
	if !strings.Contains(source(t), ":=") {
		t.Fatal("declare msg with :=")
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("variables.go")
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
