package variables

// Variables with an initializer take the value on the right.
// A short declaration, :=, declares and assigns in one step.
// https://go.dev/tour/basics/1

// Initialized returns 7 from a variable declared with var and an initializer.
func Initialized() int {
	var n int = 7
	return n
}

// Short returns "gopher" from a variable declared with :=.
func Short() string {
	msg := "gopher"
	return msg
}
