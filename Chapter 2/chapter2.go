package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to chapter 2")
	randomNum()
}

func randomNum() {
	var num = rand.Intn(10) + 1

	fmt.Println(num)
}
