package main

import "fmt"

func SwitchStatement() {
	fmt.Println("There is a cavern entrance and path to east.")
	var command = "go inside"
	switch command {
	case "go east":
		fmt.Println("You head further up to the mountains.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in the dimly lit cavern")
	case "read sign":
		fmt.Println("The sign read 'No Minors'")
	default:
		fmt.Println("Didn't quite get that")
	}

}
