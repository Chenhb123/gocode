package main

import "fmt"

func main() {
	arr := []int{1, 2, 1, 2, 3, 4, 6, 6, 4, 3, 2}
	k := 3
	// sum := maxSum(arr, k)
	sum := maxSumWin(arr, k)
	fmt.Println(sum)
}

func maxSumWin(arr []int, k int) int {
	var sum int
	if k > len(arr) {
		return -1
	}
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	for i := k; i < len(arr); i++ {
		temp := sum + arr[i] - arr[i-k]
		if temp > sum {
			sum = temp
		}
	}
	return sum
}

func maxSum(arr []int, k int) int {
	var sum int
	for i := 0; i < len(arr)-k+1; i++ {
		temp := 0
		for j := i; j < i+k; j++ {
			temp += arr[j]
		}
		if temp > sum {
			sum = temp
		}
	}
	return sum
}
