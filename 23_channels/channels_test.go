package channels_test

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/catdevman/gophlings/23_channels"
)

func TestPing(t *testing.T) {
	var got string
	err := within(time.Second, func() error {
		got = channels.Ping("hello")
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != "hello" {
		t.Fatalf("Ping(%q) = %q, want %q", "hello", got, "hello")
	}
}

func TestBufferedSum(t *testing.T) {
	code := source(t)
	if !regexp.MustCompile(`make\(\s*chan int\s*,\s*2\s*\)`).MatchString(code) {
		t.Fatal("make the channel with make(chan int, 2) so both sends fit")
	}
	var got int
	err := within(time.Second, func() error {
		got = channels.BufferedSum()
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != 3 {
		t.Fatalf("BufferedSum() = %d, want 3", got)
	}
}

func within(d time.Duration, fn func() error) error {
	errc := make(chan error, 1)
	go func() { errc <- fn() }()
	select {
	case err := <-errc:
		return err
	case <-time.After(d):
		return fmt.Errorf("timed out: a send needs a receiver, or a buffer with room")
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("channels.go")
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
