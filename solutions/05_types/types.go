package types

// Go is statically typed. These are the basic types you will use most often:
// bool, string, int, uint, float64, byte, and rune.
// https://go.dev/tour/basics/1

// Basics returns one value of each basic type listed above.
// Use true, "go", -3, uint(3), 3.14, and byte('A').
func Basics() (bool, string, int, uint, float64, byte) {
	return true, "go", -3, uint(3), 3.14, byte('A')
}
