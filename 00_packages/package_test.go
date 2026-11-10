package main_test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run .: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "Hello World!") {
		t.Fatalf("output = %q, want it to contain %q", out, "Hello World!")
	}
}
