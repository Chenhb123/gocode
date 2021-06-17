package main

import "fmt"

/*
给定一个无序的整数数组，找到其中最长上升子序列的长度
示例：
输入： [10,9,2,5,3,7,101,18]
输出：4
解释：最长上升子序列是[2,3,7,101],它的长度是4.

说明：
1.可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可
你算法的时间复杂度应该为O(n2)
进阶： 你能将算法的时间复杂度降低到O(nlogn)吗？
*/

type dp struct {
	length int
	list   []int
}

func main() {
	//length := lengthOfLIS([]int{10,9,2,5,3,7,101,18})
	//println(length)
	dp := lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Println(dp)
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			var length int
			if nums[i] > nums[j] {
				length = dp[j] + 1
			}
			if length > dp[i] {
				dp[i] = length
			}
		}
	}
	var max int
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}

func lengthOfLIS2(nums []int) dp {
	dps := make([]dp, len(nums))
	for i := 0; i < len(dps); i++ {
		dps[i].length = 1
		dps[i].list = []int{nums[i]}
	}
	for i := 0; i < len(dps); i++ {
		for j := 0; j < i; j++ {
			var length int
			var list []int
			if nums[i] > nums[j] {
				length = dps[j].length + 1
				list = append(dps[j].list, nums[i])
			}
			if length > dps[i].length {
				dps[i].length = length
				dps[i].list = list
			}
		}
	}
	var maxDp dp
	for _, v := range dps {
		if maxDp.length < v.length {
			maxDp = v
		}
	}
	return maxDp
}
