package inference

// When the right-hand side has a type, a short declaration infers it.
// You do not write the type on the left.
// https://go.dev/tour/basics/1

// Inferred returns 42, 3.14, and "go".
// Declare each variable with := and no explicit type.
func Inferred() (int, float64, string) {
	var i int
	var f float64
	var s string
	return i, f, s
}

// Kind returns the inferred type of a variable holding 42, such as "int".
func Kind() string {
	return ""
}
