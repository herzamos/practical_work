package tests

// @ requires a > 10
// @ requires a > 15 && b < 10
// @ requires a > 15 ==> b < 10
// @ requires (a > 15 ==> b < 10) || b > 15
func redundantCond2(a int, b int)