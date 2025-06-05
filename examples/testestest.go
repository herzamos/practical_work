package tests

// @ requires acc(x)
// @ ensures acc(x)
func fooooo(x *int) {
	return 
	
}
// @ assert acc(x)

func baaar() {
	var x = new(int)
	// @ assert acc(x, 1)
	fooooo(x)
}