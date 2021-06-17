package main

import "fmt"

func main() {
	// 求1亿以内质数个数
	fmt.Println(odd(10000))
}

func odd(n int) int {
	// 2 i*i <= n
	var total int
	// 100以内
	s := make([]int, n+1)
	for i := 2; i < len(s); i++ {
		target := false
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				target = true
				break
			}
		}
		if !target {
			s[i] = 1
		}
	}
	for _, v := range s {
		if v == 1 {
			total++
		}
	}
	return total
}
