package slice

// A slice is a window onto an array. The low bound defaults to 0 and the
// high bound defaults to the length. A nil slice has no array behind it.
// make creates a slice, and append grows one. Slices can hold other slices.
// https://go.dev/tour/moretypes/1

// Left returns the first two elements. Use the default low bound.
func Left(s []int) []int {
	return nil
}

// Right returns every element from index 2 onward. Use the default high bound.
func Right(s []int) []int {
	return nil
}

// Middle returns the elements at indexes 1 and 2.
func Middle(s []int) []int {
	return nil
}

// All returns every element of s by using both default bounds.
func All(s []int) []int {
	return nil
}

// NilSlice returns a nil slice. An empty slice literal is not nil.
func NilSlice() []int {
	return []int{}
}

// MakeSlice returns a slice of three zero ints with room for five.
func MakeSlice() []int {
	return nil
}

// Board returns two rows: X O, then O X.
func Board() [][]string {
	return nil
}

// Grow returns s with the extra values appended.
func Grow(s []int, vs ...int) []int {
	return s
}
