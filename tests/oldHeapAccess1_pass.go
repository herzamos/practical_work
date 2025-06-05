package tests

// @ requires acc(x)
// @ ensures acc(x) && old(*x) == *x
func oldHeapAccess1(x *int)