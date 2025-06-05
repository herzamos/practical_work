package tests

// @ requires acc(x) && acc(y)
// @ ensures *x == old(*y) && *y == old(*x)
// @ ensures acc(x) && acc(y)
func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

// @ requires acc(x) && acc(y)
// @ ensures *x == old(*y) && *y == old(*x) && acc(x) && acc(y)
func swap2(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}