package tests

type MyStruct struct {
	f int
}

//@ requires acc(&s.f)
func leakPermissions(s *MyStruct) {
  s.f = 42
}

func caller_leak() {
  s := new(MyStruct)

  leakPermissions(s)
  //@ assert acc(s)
}