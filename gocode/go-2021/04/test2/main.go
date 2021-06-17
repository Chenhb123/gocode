package main

import "fmt"

func main() {
	var list []int
	var l int
	for {
		n, _ := fmt.Scanf("%d", &l)
		if n == 0 {
			break
		} else {
			list = append(list, l)
		}
	}
	result := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		index := 0
		for j := i + 1; j < len(list); j++ {
			if list[j] > list[i] {
				index = j
				break
			}
		}
		result[i] = index
	}
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}
