package tests

// @ preserves forall k int :: {&s[k]} 0 <= k && k < len(s) ==> acc(&s[k])
// @ ensures forall k int :: {&s[k]} 0 <= k && k < len(s) ==> s[k] == old(s[k]) + n
func addToSlice(s []int, n int) {
	// @ invariant 0 <= i && i <= len(s)
	// @ invariant forall k int :: {&s[k]} 0 <= k && k < len(s) ==> acc(&s[k])
	// @ invariant forall k int :: {&s[k]} i <= k && k < len(s) ==> s[k] == old(s[k])
	// @ invariant forall k int :: {&s[k]} 0 <= k && k < i ==> s[k] == old(s[k]) + n
	for i := 0; i < len(s); i += 1 {
		s[i] = s[i] + n
	}
}

func overlappingFail() {
	s := make([]int, 3)
	l := s[:2]
	r := s[1:]
	addToSlice(l, 1)
	addToSlice(r, 1) 
}