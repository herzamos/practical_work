package tests

// @ requires acc(x, _)
// @ ensures acc(x, _)
// @ ensures y == old(*x) + 1
func foo_wildcard(x* int) (y int) {
	y = *x + 1
	return y
}


func caller_wildcard()  {
	var x *int = new(int)
	*x = 1
	foo_wildcard(x)
	*x = 2
}

