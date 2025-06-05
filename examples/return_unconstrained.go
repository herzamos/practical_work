package tests

//@ requires true
//@ ensures true
func getAnswer() int  {
  return 42
}

func caller_ret_unconstrained() {
  a := getAnswer()
  //@ assert a == 42
}