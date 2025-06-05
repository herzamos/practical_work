package tests

// @ requires acc(x, 1) && acc(y, 1/2) && acc(y, 1/2)
// @ ensures acc(x, 1/2) && acc(y, 1/2)
func leakPerm1(x *int, y *int) 