package conditional

// if can take a short statement that runs before the condition.
// else is the other branch. switch compares a value against cases.
// https://go.dev/tour/flowcontrol/1

// Sign returns "positive", "negative", or "zero".
func Sign(n int) string {
	if n > 0 {
		return "positive"
	} else if n < 0 {
		return "negative"
	} else {
		return "zero"
	}
}

// Abs returns the absolute value of n. Use else.
func Abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

// HalfIfEven returns n/2 when n is even and n otherwise.
// Declare the half with a short if statement.
func HalfIfEven(n int) int {
	if half := n / 2; n%2 == 0 {
		return half
	} else {
		return n
	}
}

// Day returns "mon", "tue", or "other". Use switch.
func Day(n int) string {
	switch n {
	case 1:
		return "mon"
	case 2:
		return "tue"
	default:
		return "other"
	}
}
