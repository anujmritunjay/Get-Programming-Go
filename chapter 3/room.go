package main

import "fmt"

func Room() {
	var room = "cave" //cave, entrance, and mountain.

	if room == "cave" {
		fmt.Println("You find yourself in dimly lit cavern")
	} else if room == "entrance" {
		fmt.Println("There is a cavern entrance here and a path to the east.")
	} else if room == "mountain" {
		fmt.Println("There is a cliff here. A path leads west down the mountain.")
	} else {
		fmt.Println("Everything is White")
	}

}
