package choose_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/catdevman/gophlings/24_select"
)

func TestFirstFromA(t *testing.T) {
	a := make(chan string)
	b := make(chan string)
	go func() {
		time.Sleep(20 * time.Millisecond)
		a <- "alpha"
	}()
	got, err := within(time.Second, func() (string, error) {
		return choose.First(a, b), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != "alpha" {
		t.Fatalf("First() = %q, want %q", got, "alpha")
	}
}

func TestFirstFromB(t *testing.T) {
	a := make(chan string)
	b := make(chan string)
	go func() {
		time.Sleep(20 * time.Millisecond)
		b <- "beta"
	}()
	got, err := within(time.Second, func() (string, error) {
		return choose.First(a, b), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != "beta" {
		t.Fatalf("First() = %q, want %q", got, "beta")
	}
}

func TestDefault(t *testing.T) {
	empty := make(chan string)
	got, err := within(time.Second, func() (string, error) {
		return choose.OrDefault(empty), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != "none" {
		t.Fatalf("OrDefault(empty) = %q, want %q", got, "none")
	}

	ready := make(chan string, 1)
	ready <- "ready"
	got, err = within(time.Second, func() (string, error) {
		return choose.OrDefault(ready), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != "ready" {
		t.Fatalf("OrDefault(ready) = %q, want %q", got, "ready")
	}
}

func within(d time.Duration, fn func() (string, error)) (string, error) {
	type result struct {
		v   string
		err error
	}
	ch := make(chan result, 1)
	go func() {
		v, err := fn()
		ch <- result{v, err}
	}()
	select {
	case r := <-ch:
		return r.v, r.err
	case <-time.After(d):
		return "", fmt.Errorf("timed out: select should take the ready case, or default")
	}
}
