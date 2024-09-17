package main

import (
	"fmt"
	"time"
)

func CountDown() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--

	}
	fmt.Println("Liftoff")
}
