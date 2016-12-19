package main

import "fmt"

var (
	message string
)

func main() {
	fmt.Println(message)
}

func init() {
	message = "Hello Go!"
}
