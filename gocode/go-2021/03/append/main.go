package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	ap2(s)
	fmt.Println(s)
	s1 := ap(s)
	fmt.Println(s, s1)
	s = ap(s)
	fmt.Println(s)
}

func ap(s []int) []int {
	s = append(s, 10)
	return s
}

func ap2(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] += 1
	}
}
