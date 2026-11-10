package inference_test

import (
	"os"
	"strings"
	"testing"

	"github.com/catdevman/gophlings/08_inference"
)

func TestInferred(t *testing.T) {
	i, f, s := inference.Inferred()
	if i != 42 || f != 3.14 || s != "go" {
		t.Fatalf("Inferred() = (%d, %v, %q), want (42, 3.14, \"go\")", i, f, s)
	}
	if strings.Contains(source(t), "var ") {
		t.Fatal("drop the explicit types and declare the variables with :=")
	}
}

func TestKind(t *testing.T) {
	if got := inference.Kind(); got != "int" {
		t.Fatalf("Kind() = %q, want %q", got, "int")
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("inference.go")
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
