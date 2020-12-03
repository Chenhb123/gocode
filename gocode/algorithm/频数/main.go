package main

import "fmt"

/*
给定一个非空且只包含非负数的整数数组 nums, 数组的度的定义是指数组里任一元素出现频数的最大值。
你的任务是找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
*/

func main() {
	nums := []int{1, 2, 2, 3, 1}
	result := findShortestSubArray(nums)
	fmt.Println("result:", result)
}

func findShortestSubArray(nums []int) int {
	// 确定nums的度
	numap := make(map[int]int, 0)
	for _, v := range nums {
		numap[v] = 0
	}
	for _, v := range nums {
		if _, ok := numap[v]; ok {
			numap[v]++
		}
	}
	var rec int
	for _, v := range numap {
		if rec < v {
			rec = v
		}
	}
	fmt.Println("rec:", rec)
	fmt.Println("numap:", numap)
	// 获取全部连续子数组
	var aims []int
	for i, v := range numap {
		if v == rec {
			aims = append(aims, i)
		}
	}
	fmt.Println("aims:", aims)
	start, end := 0, len(nums)-1
	var result, length int
	result = len(nums)
	length = len(nums)
	for start < end {
		index := start + 1
		temp := nums[index : end+1]
		var t int
		starget, etarget := false, false
		for _, v := range temp {
			for _, a := range aims {
				if a == v {
					t++
				}
			}
			if t == rec {
				// 满足要求
				starget = true
				if length < len(temp) {
					length = len(temp)
				}
				break
			}
		}
		temp = nums[start:end]
		t = 0
		for _, v := range temp {
			for _, a := range aims {
				if a == v {
					t++
				}
			}
			if t == rec {
				etarget = true
				if length < len(temp) {
					length = len(temp)
				}
				break
			}
		}
		fmt.Println("starget, etarget, length:", starget, etarget, length)
		if starget {
			start++
		}
		if etarget {
			end--
		}
		fmt.Println("start, end:", start, end)
		if starget || etarget {
			if length < result {
				result = length
			}
		}
		if !starget && !etarget {
			break
		}
	}
	return result
}
