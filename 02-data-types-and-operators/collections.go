package main

import "fmt"

func main() {
	arrays()
	initializedArrays()
	arrayLength()
	slices()
	initializingSlices()
	largeSlices()
	maps()
}

func arrays() {
	someArray := [3]int{}

	someArray[0] = 16
	someArray[1] = 32
	someArray[2] = 64

	fmt.Println(someArray)
}

func initializedArrays() {
	initializedArray := [...]int{12, 21, 42}
	fmt.Println(initializedArray)
}

func arrayLength() {
	someArray := [...]int{21, 112, 4124}
	fmt.Println(len(someArray))
}

// Subset of an array - go takes care of sizing / resizing the array.
func slices() {
	someArray := [...]int{1, 2, 3}
	someSlice := someArray[:]

	fmt.Println(someSlice)

	someSlice = append(someSlice, 4)

	fmt.Println(someSlice)
}

// Not very efficient as the underlying array is always exactly the same size as the slice.
func initializingSlices() {
	someSlice := []string{"Hello", "World"}
	fmt.Println(someSlice)
	fmt.Println(len(someSlice))
}

func largeSlices() {
	bigSlice := make([]int, 100)

	bigSlice[0] = 21
	bigSlice[1] = 42

	fmt.Println(bigSlice)
	fmt.Println(len(bigSlice))
}

func maps() {
	someMap := make(map[string]int)
	fmt.Println(someMap)

	someMap["hello"] = 21
	someMap["world"] = 42

	fmt.Println(someMap)
	fmt.Println(someMap["world"])
	fmt.Println(someMap["does not exist"])
}
