package tests

//@ requires x < 43 && x > 41
func onlyAccepts42(x int) {
  _ = x + 1
}

func caller_one_input() {
  onlyAccepts42(42)     
  // onlyAccepts42(7)   
}