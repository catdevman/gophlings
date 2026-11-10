package arrays

// An array's length is part of its type. [3]int and [4]int are different types.
// Arrays are values, so they can be compared with ==.
// https://go.dev/tour/moretypes/1

// Evens returns the array [2, 4, 6].
func Evens() [3]int {
	var a [3]int
	return a
}

// Last returns the last element of a.
func Last(a [4]string) string {
	return ""
}

// Same reports whether a and b hold the same elements.
func Same(a, b [2]int) bool {
	return false
}
