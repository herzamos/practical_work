package tests

type Number interface {
	N() (n int)
}

func client_nil_interface() {
	var n Number = (*PositiveNumber)(nil)
	// @ assert n != nil
	// @ assert typeOf(n) == type[*PositiveNumber]
	n.N()
}

type PositiveNumber struct{ x int }

func (s *PositiveNumber) N() (n int) {
	// @ assert s != nil // error
	n = s.x
	return n
}
