package tests


var global /*@ @ @*/  int = 40

// @ requires acc(x)
// @ requires x == &global
func ff(x *int) {
}


func random() (x *int) 

func refute(x int, y int) {
	// @ inhale x > 0
	// @ inhale y > 0
	// @ refute x == y
}