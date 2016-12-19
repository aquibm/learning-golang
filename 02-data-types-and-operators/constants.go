package main

// Constants can be defined in-line or in a variable block
const (
	zeroth = "Hello world!"
)

// Constants without assignments are given an integer value starting at 0
const (
	first = iota
	second
	third
)

// Constant-expression: Each constant defined in the block is run through 1 << (10 * iota)
const (
	one = 1 << (10 * iota)
	two
	three
)

func main() {
	println(zeroth)

	println(first)
	println(second)
	println(third)

	println(one)
	println(two)
	println(three)
}
