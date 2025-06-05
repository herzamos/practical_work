package tests

type MyStruct2 struct {
  x int
  y int
}

//@ requires acc(&s.x) && acc(&s.y)
//@ ensures acc(&s.x) && acc(&s.y) && s.x == old(s.x) + 1
func updateBoth(s *MyStruct2) {
  s.x = s.x + 1
  s.y = s.y + 1 // Unconstrained!
}

//@ requires acc(s)
func caller_heap_unconstrained(s *MyStruct2) {
	s.x = 0
	s.y = 0
	updateBoth(s)
	//@ assert s.x == 1 && s.y == 1
}