package main

import (
	"fmt"
)

// Exercise 1

// Write a function that takes two integers and returns the larger one.

func largeInt(a, b int) int {

	if a >= b {
		return a
	}
	return b

}

func main() {
	result := largeInt(6, 6)
	fmt.Println(result)

	fmt.Println(isEven(8))
	fmt.Println(isEven(11))

	fmt.Println(squareIt(6))

	resultme := multiply(2, 3)

	fmt.Println(resultme)

	result5 := largest(2, 9, 3, 17, 8)
	fmt.Println(result5)

	fmt.Println(vAverage(10, 20, 30))

	// Exercise 7

	// Store an anonymous function inside a variable.

	// It should print

	// Learning Go

	// Then call it.

	pntmsg := func() {
		fmt.Println("Learning Go")
	}
	pntmsg()

	mycounter := counter()

	fmt.Println(mycounter())
	fmt.Println(mycounter())
	fmt.Println(mycounter())

	mycounter2 := counter2()

	fmt.Println(mycounter2())
	fmt.Println(mycounter2())
	fmt.Println(mycounter2())

	mycounter3 := decreasedCounter()
	fmt.Println(mycounter3())
	fmt.Println(mycounter3())
	fmt.Println(mycounter3())

	(printup(5))

	fmt.Println(sum(5))

	fmt.Println(power(2, 5))

	fmt.Println(Reverse("Hello thr"))

}

// Write a function called isEven()

func isEven(num int) bool {

	if num%2 == 0 {
		return true
	}
	return false
}

// Write a function that returns the square of a number.

func squareIt(num int) int {
	value := num * num
	return value
}

// Write your own variadic function named multiply.

func multiply(nums ...int) int {
	multiplied := 1
	for _, num := range nums {
		multiplied *= num
	}
	return multiplied

}

func largest(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max

}

// Create a variadic function that returns the average.

func vAverage(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	total := 0.0
	for _, num := range nums {
		total += num

	}
	return total / float64(len(nums))
}

//  Exercise 9

// Create a closure that starts at 100.

func counter() func() int {
	i := 100
	return func() int {
		i++
		return i
	}
}

// Modify the closure so it increases by 5 every time.
func counter2() func() int {
	i := 100
	return func() int {
		i += 5
		return i
	}
}

// Exercise 11 (Challenge)

//Make a closure that decreases instead.

func decreasedCounter() func() int {
	i := 50
	return func() int {
		current := i
		i--
		return current
	}

}

// 	Exercise 12

// Print numbers from 1 to n using recursion.
func printup(n int) {
	if n <= 0 {
		return
	}
	printup(n - 1)

	fmt.Println(n)
}

// Exercise 13

// Write a recursive function that adds all numbers from 1 to n.

func sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sum(n-1)
}

// Exercise 14

// Write a recursive function that calculates powers.
func power(n int, raise int) int {
	if n == 0 {
		return 0

	}
	if raise == 0 {
		return 1
	}
	return n * power(n, raise-1)
}

func Reverse(s string) string {
	if len(s) <= 1 {
		return s
	}
	return Reverse(s[1:]) + string(s[0])
}
