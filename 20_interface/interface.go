package shapes

// An interface is a set of method signatures. A type implements it by
// having the methods. There is no implements keyword.
// A nil interface has no type and no value. An interface can also hold
// a typed nil, and that value is not nil.
// any is the empty interface and can hold a value of any type.
// A type alias, written with =, is another name for the same type.
// https://go.dev/tour/methods/1

type Celsius float64

// Fahrenheit should be an alias for float64, not a new type.
type Fahrenheit float64

func AsFloat(f Fahrenheit) float64 {
	var out float64 = f
	return out
}

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return 0
}

// Vertex implements Abser because it has Abs. Nothing else is required.
var _ Abser = Vertex{}

func Measure(a Abser) float64 {
	return a.Abs()
}

type Speaker interface {
	Speak() string
}

type Quiet struct{}

func (q *Quiet) Speak() string {
	return ""
}

// NilSpeaker returns a Speaker whose dynamic value is a nil *Quiet.
func NilSpeaker() Speaker {
	return nil
}

// MissingSpeaker returns a nil Speaker, with no dynamic type.
func MissingSpeaker() Speaker {
	var s Speaker
	return s
}

// Pack stores v in an empty interface and returns it.
func Pack(v any) any {
	return nil
}
