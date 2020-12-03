package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 3)
	copy(b, a)
	fmt.Println(b)
	test(b...)
	fmt.Println(b)
}

func test(a ...int) {
	for i := range a {
		a[i] += 100
	}
}
