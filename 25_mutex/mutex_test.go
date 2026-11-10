package mutexes_test

import (
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/catdevman/gophlings/25_mutex"
)

func TestCounter(t *testing.T) {
	if !strings.Contains(source(t), "sync.Mutex") {
		t.Fatal("protect the counter with a sync.Mutex")
	}
	var c mutexes.Counter
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Add(1)
			}
		}()
	}
	wg.Wait()
	if got := c.Value(); got != 3200 {
		t.Fatalf("Value() = %d, want 3200", got)
	}
}

func source(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("mutex.go")
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
