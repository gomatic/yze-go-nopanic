package a

// boom calls the built-in panic in a non-test file and must be flagged.
func boom() {
	panic("x") // want `panic is not permitted`
}

// ok calls a different built-in (len) and must NOT be flagged.
func ok() {
	_ = len("x")
}

// shadow calls a local variable named panic that shadows the built-in, so the
// callee is not the predeclared panic and must NOT be flagged.
func shadow() {
	panic := func(string) {}
	panic("y")
}
