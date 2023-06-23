package main

import "fmt"

func main() {

	var carName = "Varena"

	fmt.Println("This folder is for the variable and constant")
	fmt.Print("This folder is for the variable and constant \n")

	carName = "Maruti"

	const ruppe = 1
	fmt.Printf("The car name is %v \n", carName)
	// ruppe = 5 -> this is wrong cuz we can't assign the variable to the new one.
	fmt.Printf("I have %v coin\n", ruppe)

	//  Variable using the data type like in c/c++, java etc.

	var name string = "vikash"
	fmt.Println("Hey my name is \n", name)

	// OR
	Greet := "Good MOrning"

	fmt.Printf("Hey %v %v, How are you \n", name, Greet)

	// Integer

	var price int = 50
	fmt.Printf("The price is %v", price)
}
