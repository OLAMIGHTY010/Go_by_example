package main

import "fmt"




func countdown (number int){
	fmt.Println(number)

	if number == 0{
		return
	}
	countdown(number-1)
}

func main (){

	fmt.Println(isPalindrome("racecar"))
}



func countdown2 (nums int){
	if nums == 0{
		fmt.Println("this is a blast")
		return
	}
	fmt.Println(nums)
	countdown2(nums - 1)
	fmt.Println("recurring the numbers", nums)
}


func isPalindrome(s string)bool{
	if len(s)<=1 {
		return true
	}
	first := s[0]
	last := s[len(s)-1]

	if first != last{
		return false
	}
	middle := s[1: len(s)-1]
	return isPalindrome(middle)

}