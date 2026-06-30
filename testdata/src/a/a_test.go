package a

// helper calls the built-in panic in a _test.go file, which is allowed, so it
// must NOT be flagged (exercises the test-file-skip branch).
func helper() {
	panic("allowed in tests")
}
