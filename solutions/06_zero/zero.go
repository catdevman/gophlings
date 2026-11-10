package zero

// A variable declared without an initializer gets the zero value of its type:
// 0 for numbers, false for booleans, "" for strings, and nil for slices.
// https://go.dev/tour/basics/1

// Zeros returns the zero value of an int, a float64, a string, a bool, and a []int.
// Declare each one with var and no initializer.
func Zeros() (int, float64, string, bool, []int) {
	var i int
	var f float64
	var s string
	var b bool
	var items []int
	return i, f, s, b, items
}
