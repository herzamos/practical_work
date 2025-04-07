package tests

type MyStruct struct {
	f int
}

//@ requires acc(&s.f)
//@ ensures acc(&s.f)
func onlyReadF(s *MyStruct) (f int) {
  f = s.f 
  return f
}

//@ requires acc(s, _)
func caller_only_reads(s *MyStruct) {
   
  var f = onlyReadF(s)           
}