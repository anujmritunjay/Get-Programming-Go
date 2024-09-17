package main

import "fmt"

func If() {
	var command = "go east"
	if command == "go east" {
		fmt.Println("You head further up to mountains")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out rest of the your life")
	} else {
		fmt.Println("Didn't quite get that")
	}
}
