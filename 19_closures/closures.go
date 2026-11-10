package funcs

import "net/http"

// Functions are values. A closure is a function that uses variables
// from the surrounding function. http.HandlerFunc is a function type
// used as an HTTP handler. https://go.dev/tour/moretypes/1

// Adder returns a function that adds its argument to a running sum
// and returns the new sum. Each Adder keeps its own sum.
func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		return sum
	}
}

// Hello writes the body "Hello, gopher".
func Hello(w http.ResponseWriter, r *http.Request) {
}

// Route returns a mux that serves Hello at /hello.
func Route() *http.ServeMux {
	return http.NewServeMux()
}
