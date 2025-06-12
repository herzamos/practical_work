package tests

func f(x *int) {
	// @ inhale acc(x)
	// @ exhale forperm x: Ref [x] :: false
}