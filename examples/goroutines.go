package tests

type Counter struct {
	count int
}

// @ requires acc(&c.count)
// @ ensures acc(&c.count)
func (c *Counter) Increment() {
	c.count += 1
}

func main() {
	ctr := new(Counter)
	go ctr.Increment()
	go ctr.Increment() // error
}