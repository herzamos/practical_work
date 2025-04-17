package tests

/*@
pred P(x int) {
	true
}
@*/

// @ ensures P(res)
func m() (res int) {
	// @ fold P(res)
	return 10
}