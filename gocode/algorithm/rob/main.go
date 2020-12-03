package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 5, 8, 6, 4, 3, 2, 4, 6}
	count := rob(nums)
	fmt.Println(count)
}

// nums 1 3 2 5  8  6  4  3  2  4  6
// robs 1 3 3 8 11 14 15 17 17 21 23
// robs[i] = max(robs[i-2]+nums[i], robs[i-1])
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var pre, res int
	robs := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			pre = nums[0]
			robs[0] = nums[0]
			continue
		}
		if i == 1 {
			robs[1] = nums[1]
			res = nums[1]
			if nums[0] > nums[1] {
				res = nums[0]
				robs[1] = nums[0]
			}
			continue
		}
		robs[i] = robs[i-1]
		if pre+nums[i] > res {
			pre, res = res, pre+nums[i]
		} else {
			pre = res
		}
		if robs[i-2]+nums[i] > robs[i-1] {
			robs[i] = robs[i-2] + nums[i]
		}
	}
	fmt.Println(res)
	return robs[len(robs)-1]
}
