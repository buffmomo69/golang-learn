package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointer")

	// var pointer *int
	// fmt.Print(pointer)

	myNumber := 23
	var myNumberPointer = &myNumber

	fmt.Print(myNumberPointer)
	fmt.Print(*myNumberPointer)

	*myNumberPointer = *myNumberPointer + 2
	fmt.Println("New v alue is", myNumber)
}
