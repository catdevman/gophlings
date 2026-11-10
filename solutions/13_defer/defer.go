package deferred

import "strconv"

// A deferred call runs when the surrounding function returns.
// Deferred calls run in last-in, first-out order.
// A named result lets a deferred call add to the value that is returned.
// https://go.dev/tour/flowcontrol/1

// Stacked returns "open", "write", "close", "flush", in that order.
// Register the close and flush calls with defer so flush runs last.
func Stacked() (out []string) {
	out = append(out, "open")
	defer func() { out = append(out, "flush") }()
	defer func() { out = append(out, "close") }()
	out = append(out, "write")
	return
}

// Countdown returns "321".
// Defer a call for each number from 1 through 3, and pass the number
// into the deferred function so it keeps the value from that iteration.
func Countdown() (s string) {
	for i := 1; i <= 3; i++ {
		defer func(n int) {
			s += strconv.Itoa(n)
		}(i)
	}
	return
}
