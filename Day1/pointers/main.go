package main



import "fmt"

func change(x *int) {
	*x = *x + 5
}

func main() {
	a := 20

	change(&a)

	fmt.Println(a)
}