package numeric

// Numeric constants are high precision values. They do not overflow
// until they are given a type that cannot hold them.
//
// Define Big as 1 shifted left by 100, and Small as Big shifted right by 99.
// Small fits in an int. Big does not, so only pass Big to NeedFloat.
// https://go.dev/tour/basics/1

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func NeedInt(x int) int { return x*10 + 1 }

func NeedFloat(x float64) float64 { return x * 0.1 }

func UseSmallInt() int { return NeedInt(Small) }

func UseSmallFloat() float64 { return NeedFloat(Small) }

func UseBigFloat() float64 { return NeedFloat(Big) }
