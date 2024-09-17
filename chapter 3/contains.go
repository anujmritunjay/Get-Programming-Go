package main

import (
	"fmt"
	"strings"
)

func Contains() {
	fmt.Println("You find yourself in dimly lit cavern")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You can leave the cave: ", exit)
}
