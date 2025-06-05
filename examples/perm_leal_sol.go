package tests

// @ requires acc(x)
// @ ensures acc(x, _)
func leakingOrNot(x *int ) {
	return
}

func testLeak() {
	var x = new(int)
	// @ assert acc(x)
	leakingOrNot(x)
	// @ assert acc(x)
}