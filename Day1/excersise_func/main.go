package main

import (
	"fmt"
)

// 1. Minimum
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//2. absolur number

func abs(n int) int {
	if n < 0 {
		return n * (-1)
	}
	return n
}

func isLeap(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}
func celsiusToFahrenheit(c float64) float64 {
	F := float64((c * 1.8) + 32)

	return float64(F)
}
func countPositive(nums []int) int {
	count := 0
	for _, num := range nums {
		if num > 0 {
			count++
		}
	}
	return count
}

// Find the smallest number in a slice.

func smallest(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num

		}
	}
	return min

}

// Return the sum of only the even numbers.
func evenSum(nums []int) int {
	total := 0
	for _, num := range nums {
		if num%2 == 0 {
			total += num
		}
	}
	return total
}
// Return a new slice containing only odd numbers.
func oddNumbers(nums []int) []int{
	slice := make([]int, 0, len(nums)/2)

	for num := range nums{
		if num%2 !=0{
			slice = append(slice,num)
			
		}
	}
	return slice
	
}

// Return how many numbers are greater than 100.

func greaterThan100(nums ...int) int {
	count := 0
	for _, num := range nums{
		if num > 100{
			count++
		}
	}
	return count

}
// Return the product of only the even numbers.
func evenProduct(nums ...int)int{
	total := 1
	for _, num := range nums{
		if num%2 == 0{
			total *= num
		}
	}
	return total
}



func main() {
	fmt.Println(min(3, 4))

	fmt.Println(abs(-4))

	fmt.Println(isLeap(2023))
	fmt.Println(isLeap(2024))
	fmt.Println(isLeap(2020))

	fmt.Println(celsiusToFahrenheit(0))

	s := []int{6, 7, 4, 5, 1, 8}
	fmt.Println(countPositive(s))
	fmt.Println(smallest(s))
	fmt.Println(evenSum(s))
	fmt.Println(oddNumbers(s))
	fmt.Println(small(10, 5, 20))
	fmt.Println(greaterThan100(101,102,103,100,99))
	fmt.Println(evenProduct(1,2,3,4,5,6))

	next := fibonacci()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	




}

func small(nums ...int) int{
	Minimum := nums[0]
	for _, num := range nums[1:] {
		if num < Minimum{
			Minimum = num
		}
		
	}
	return Minimum

}

// Create a closure that starts at 1000 and decreases by 100 every call.
func decreases() func() int{
	starts := 1000
	return func() int {
		starts -= 100
		return starts
	}
}

func decreases1() func ()int{
	start := 10
	return func() int{
		if start <= 0{
			return 0
		}
		start -=10 
		return start
	}
}
// Create a closure that doubles every call.
func doubles() func ()int{
	i := 10
	return func() int {
		i*=2
		return i
	}
}
func fibonacci() func() int{
	a:= 0
	b:= 1
	return  func() int {
		current := a
		a,b = b, a+b
		return current
	}

}