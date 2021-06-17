package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2}
	var result []int
	for _, v := range s1 {
		target := false
		for _, t := range s2 {
			if t == v {
				target = true
				break
			}
		}
		if !target {
			result = append(result, v)
		}
	}
	fmt.Println(result)
}
