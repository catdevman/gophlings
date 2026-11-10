package goroutine_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/catdevman/gophlings/22_goroutine"
)

func TestSum(t *testing.T) {
	var got int
	err := within(time.Second, func() error {
		got = goroutine.Sum([]int{1, 2, 3})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != 6 {
		t.Fatalf("Sum() = %d, want 6", got)
	}
}

func within(d time.Duration, fn func() error) error {
	errc := make(chan error, 1)
	go func() { errc <- fn() }()
	select {
	case err := <-errc:
		return err
	case <-time.After(d):
		return fmt.Errorf("timed out: the goroutine has to send the sum")
	}
}
