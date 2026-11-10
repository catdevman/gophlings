package deferred

// A deferred call runs when the surrounding function returns.
// Deferred calls run in last-in, first-out order.
// A named result lets a deferred call add to the value that is returned.
// https://go.dev/tour/flowcontrol/1

// Stacked returns "open", "write", "close", "flush", in that order.
// Register the close and flush calls with defer so flush runs last.
func Stacked() []string {
	var out []string
	out = append(out, "open")
	out = append(out, "flush")
	out = append(out, "close")
	out = append(out, "write")
	return out
}

// Countdown returns "321".
// Defer a call for each number from 1 through 3, and pass the number
// into the deferred function so it keeps the value from that iteration.
func Countdown() string {
	return ""
}
