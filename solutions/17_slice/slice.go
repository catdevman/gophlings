package slice

// A slice is a window onto an array. The low bound defaults to 0 and the
// high bound defaults to the length. A nil slice has no array behind it.
// make creates a slice, and append grows one. Slices can hold other slices.
// https://go.dev/tour/moretypes/1

// Left returns the first two elements. Use the default low bound.
func Left(s []int) []int {
	return s[:2]
}

// Right returns every element from index 2 onward. Use the default high bound.
func Right(s []int) []int {
	return s[2:]
}

// Middle returns the elements at indexes 1 and 2.
func Middle(s []int) []int {
	return s[1:3]
}

// All returns every element of s by using both default bounds.
func All(s []int) []int {
	return s[:]
}

// NilSlice returns a nil slice. An empty slice literal is not nil.
func NilSlice() []int {
	var s []int
	return s
}

// MakeSlice returns a slice of three zero ints with room for five.
func MakeSlice() []int {
	return make([]int, 3, 5)
}

// Board returns two rows: X O, then O X.
func Board() [][]string {
	return [][]string{
		{"X", "O"},
		{"O", "X"},
	}
}

// Grow returns s with the extra values appended.
func Grow(s []int, vs ...int) []int {
	return append(s, vs...)
}
