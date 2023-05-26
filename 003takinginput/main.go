package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	//capital Public lower local
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for out pizza:")
	
	//comma ok  || error ok syntax
	//underscore to discard the error value

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T", input)

}
