package main

import "fmt"

func DaysToAlphaCent() {
	const lightSpeed = 299792
	const secondPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha centauri is ", distance, " km away.")

	days := distance / lightSpeed / secondPerDay

	fmt.Println("That is ", days, "days of travel at light speed.")
}
