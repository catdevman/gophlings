package functions_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/catdevman/gophlings/01_functions"
)

func TestFoo(t *testing.T) {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	functions.Foo()
	w.Close()
	os.Stdout = old
	io.Copy(&buf, r)

	out := buf.String()
	if !strings.Contains(out, "Important parameter") {
		t.Fatalf("Foo should call the unexported function with a parameter\noutput:\n%s", out)
	}
}
