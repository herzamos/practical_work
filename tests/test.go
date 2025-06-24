package tests

// @ requires x == 1/2 && acc(y, x) && acc(y, 1/3)
// @ requires acc(z, 1/2)
// @ requires acc(a, 1/2)
// @ ensures acc(a, 1/2)
func ff(x int, y *int, z *int, a *int) {
	// @ ghost var x_leakCheck int = x
}