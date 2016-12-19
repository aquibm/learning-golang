package main

func main() {
	var someInteger int
	someInteger = 1337
	println(someInteger)

	var someFloat32 float32 = 21.
	println(someFloat32)

	someString := "Hello, world!"
	println(someString)

	someComplex := complex(2, 3)
	println(someComplex)
	println(real(someComplex))
	println(imag(someComplex))
}
