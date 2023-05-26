package main

import "fmt"

func fu() {

	fmt.Print(add(1, 3))

}

func add(a int, b int) int {

	return a + b
}
