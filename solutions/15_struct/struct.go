package vertex

// A struct groups fields. A pointer receiver can change the value it is called on.
// A value receiver sees a copy. Go turns v.X into (*v).X when v is a pointer.
// https://go.dev/tour/moretypes/1

type Vertex struct {
	X int
	Y int
}

// New returns a Vertex with the given field values.
func New(x, y int) Vertex {
	return Vertex{X: x, Y: y}
}

// Move adds dx and dy to the vertex p points at.
func Move(p *Vertex, dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Sum returns X + Y. Use a value receiver.
func (v Vertex) Sum() int {
	return v.X + v.Y
}

// Scale multiplies both fields by f. Use a pointer receiver so the caller changes.
func (v *Vertex) Scale(f int) {
	v.X *= f
	v.Y *= f
}
