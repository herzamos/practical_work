package tests

// @ requires p > 10 && acc(x, 1/p)
// @ ensures q > 0 && acc(x, 1/q)
func leakPerm3(x *int, p int) (q int) {
	q = p + 1
	// @ assert p > q
	return q
}