package pointers

// A pointer holds the address of a value. & takes an address. * follows one.
// https://go.dev/tour/moretypes/1

// Set stores v at the address p.
func Set(p *int, v int) {
	*p = v
}

// Get returns the int stored at p.
func Get(p *int) int {
	return *p
}

// NewInt returns a pointer to a new int holding v.
func NewInt(v int) *int {
	n := v
	return &n
}
