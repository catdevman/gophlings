package choose

// select waits on several channel operations and runs the one that is ready.
// default runs when no other case is ready.
// https://go.dev/tour/concurrency/1

// First returns the first value received from either a or b.
func First(a, b <-chan string) string {
	select {
	case v := <-a:
		return v
	case v := <-b:
		return v
	}
}

// OrDefault returns a value from c, or "none" when c has nothing waiting.
func OrDefault(c <-chan string) string {
	select {
	case v := <-c:
		return v
	default:
		return "none"
	}
}
