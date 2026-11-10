package constant

// A constant is declared with const. It cannot be changed with =.
// A typed constant has an explicit type. An untyped constant takes a type when it is used.
// https://go.dev/tour/basics/1

const Title = "gophlings"

const Answer int = 42

// App returns the program title.
func App() string {
	return Title
}

// Meaning returns Answer.
func Meaning() int {
	return Answer
}
