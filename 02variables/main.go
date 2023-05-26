package main

import "fmt"

func main() {

	var username string = "name"
	var isLoggedIn bool = true
	var smallInt uint8 = 255
	//5 values after decimal
	var smallFloat float32 = 2.12

	//more
	var largeFLoat float64 = 2.121312313131312313131313

	fmt.Printf("Variable is of type : %T \n", username)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", smallInt)
	fmt.Printf("Variable is of type : %T \n", smallFloat)
	fmt.Printf("Variable is of type : %T \n", largeFLoat)

	//default values and aliass
	var anotherVariable int
	var anotherString string

	fmt.Println(anotherVariable)
	fmt.Println(anotherString)

	//implicit of
	var websitedf = "fuckme"

	numberOfUsers := 1

	fmt.Print(websitedf, numberOfUsers)

}
