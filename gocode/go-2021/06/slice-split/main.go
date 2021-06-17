package main

import (
	"fmt"
)

func main() {
	a := make([]int, 20)
	a = []int{7, 8, 9, 10}
	b := a[15:16]
	fmt.Println(b)
}
