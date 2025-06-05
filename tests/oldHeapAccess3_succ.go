package tests

/*@
pred elements(l *List) {
	acc(&l.Value) && acc(&l.next) && (l.next != nil ==> elements(l.next))
}
@*/