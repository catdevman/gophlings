package deferred_test

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/catdevman/gophlings/13_defer"
)

func TestStacked(t *testing.T) {
	got := deferred.Stacked()
	want := []string{"open", "write", "close", "flush"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Stacked() = %#v, want %#v", got, want)
	}
}

func TestCountdown(t *testing.T) {
	if got := deferred.Countdown(); got != "321" {
		t.Fatalf("Countdown() = %q, want %q", got, "321")
	}
}

func TestUsesDefer(t *testing.T) {
	if strings.Count(source(t), "defer ") < 2 {
		t.Fatal("use defer for the stacked calls and the countdown")
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("defer.go")
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
