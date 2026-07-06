package main

import (
	"fmt"
	"time"
)

func main() {
	hours := 222
	switch {
	case hours >= 0 && hours <= 7:
		fmt.Println("THis is midnight shift")
	case hours >= 7 && hours <= 17:
		fmt.Println("day shift")
	case hours >= 17 && hours <= 24:
		fmt.Println("night shift")
	default:
		fmt.Println("invalid timing and shift")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Bro this is weekend, I dont work")
	case time.Now().Weekday():
		fmt.Println("yo! im active for learning and working ")
	default:

	}

}
