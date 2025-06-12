package tests

// @ ensures a > b ⇒ c == a && b > a ⇒ c == b
func max(a int, b int) (c int) {
    if a > b  { 
		c  = a; 
		return c 
	} else if b > a {
		 c = b; 
		 return b 
	} else { 
		c = 100; 
		return c 
	}
}