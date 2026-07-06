package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println(s, s == nil, len(s) == 0)
	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println(s[2])
	s = append(s, "d", "e", "f")

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	l := s[:5]
	fmt.Println("slice1", l)
	l = s[2:5]
	fmt.Println("slie2", l)
	l = s[2:]
	fmt.Println("slice3:", l)
	t := []string{"g", "h", "i"}
	fmt.Println(t)
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t==t2")
	}
	twoD := make([][]int, 12)
	for a := range 12 {
		innerLen := a + 1
		twoD[a] = make([]int, innerLen)
		for b := range innerLen {
			twoD[a][b] = a + b

		}
	}
	fmt.Println(twoD)

	m := make(map[string]int)
	m ["number"] = 99
	m["text"] = 21
	fmt.Println(m)

	v1 := m["text"]
	v2 := m["number"]

	fmt.Println(v1, v2)
	fmt.Println(len(m))

}
