package tests

type MyStruct struct {
	f int
}

//@ requires acc(&s.f, _)
//@ ensures acc(&s.f, _)
func noReadF(s *MyStruct) {}

//@ requires acc(s, _)
func caller_no_reads(s *MyStruct) {
  noReadF(s)           
}