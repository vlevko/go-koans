package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

const two = 2

func aboutPanics() {
	// assert(__delete_me__) // panics are exceptional errors at runtime

	n := divideFourBy(two)
	assert(n == 2) // panics are exceptional errors at runtime
}
