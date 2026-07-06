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
	countdown(3)
	countdown2(4)
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