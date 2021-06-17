package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(dp(n))
}

func dp(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	result := make([]int, n)
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			result[i] = result[i/2] + 1
		} else {
			result[i] = result[(i+1)/2] + 2
			if result[i] > result[(i-1)/2]+2 {
				result[i] = result[(i-1)/2] + 2
			}
		}
	}
	return result[n]
}
