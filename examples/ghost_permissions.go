package tests

const M = 42

type MyStruct struct {
	f int
}

/*@
ghost
requires acc(&s.f)
decreases
pure func isZero(s *MyStruct) bool {
    return s.f == 0
}
@*/

//@ requires acc(s)
func caller_ghost_read_permissions(s *MyStruct) {
  // @ assert acc(s)
  s.f = 0
  // @ assert isZero(s)
  // @ assert acc(s)  
}