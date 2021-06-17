package main

import "fmt"

func main() {
	//s := []int{5, -3, 4, -2, 1, -5, 3, -4, 2, -1, 0}
	s := []int{-1, 2, -3, 2, -2, 3, -4, 5, 3}
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}

func quickSort(s []int, left, right int) {
	if left >= right {
		return
	}
	// 调整基准的索引
	j := quickAdjust(s, left, right)
	// 调整左边的序列
	quickSort(s, left, j-1)
	// 调整右边的序列
	quickSort(s, j+1, right)
}

// 返回调整后基准数的位置，基准数取s[right]
func quickAdjust(s []int, left, right int) int {
	i, j := left, right
	for i < j {
		fmt.Println("start, i,j:", i, j)
		// 从左向右找大于基准的数
		for i < j && s[i] <= s[right] {
			i++
		}
		// 从右向左找小于基准的数
		for i < j && s[j] >= s[right] {
			j--
		}
		// 交换
		fmt.Println("end,i, j :", i, j)
		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
	// 调整基准的位置
	s[right], s[j] = s[j], s[right]
	return j
}
