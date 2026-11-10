package functions

import (
	"errors"
	"fmt"
)

// Foo - change this function name from foo to Foo, Uppercase function names are for allowing calls from outside through the package.
func foo() {
	fmt.Println("Change this function name to Foo if you haven't already")
	// since bar is within the same file and therefore the same package you can call it without the need to import a package or using uppercase
	bar("Important parameter")
}

// bar - is a function that can only be called from within the "functions" package, is also takes a single parameter "msg" which is a string
func bar(msg string) {
	fmt.Println("This can be called from within the \"functions\" package", msg)
}

// In Go you can use multiple returns, a common pattern is return values first and to have the last item that is returned to be an error
func baz() (int, error) {
	if 4%2 == 0 {
		return 42, nil
	} else {
		return 0, errors.New("error from baz")
	}
}
