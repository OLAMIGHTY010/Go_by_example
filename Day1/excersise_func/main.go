package main

import "fmt"

// 1. Minimum
func min (a,b int)int{
	if a < b{
		return a
	}
	return b
}

//2. absolur number

func abs(n int)int{
	if n < 0{
		return n*(-1)
	}
	return n
}

func isLeap(year int) bool {
	if year%4 == 0 && year%100 !=0 || year %400 == 0{
		return true
	}
	return false
}
func celsiusToFahrenheit(c float64) float64{
	F := float64((c * 1.8) + 32)

	return float64(F)
}
func countPositive(nums []int) int{
	count := 0
	for _, num := range nums{
		if num > 0{
			count ++
		}
	}
	return count
}

// Find the smallest number in a slice.

func smallest(nums []int) int{
	min := nums[0]
	for _, num  := range nums{
		if num < min {
			min = num

		}
	}
	return min

}

func main (){
	fmt.Println(min(3,4))

	fmt.Println(abs(-4))

	fmt.Println(isLeap(2023))
	fmt.Println(isLeap(2024))
	fmt.Println(isLeap(2020))

	fmt.Println(celsiusToFahrenheit(0))

	s := []int{3, 7, 0, 5, 1, 8}
	fmt.Println(countPositive(s))
	fmt.Println(smallest(s))
	


}
