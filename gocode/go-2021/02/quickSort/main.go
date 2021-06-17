package main

import "fmt"

// 快速排序

func main() {
	s := []int{-1, 2, -3, 2, -2, 3, -4, 5, 3}
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}

func quickSort(s []int, left, right int) {
	if left >= right {
		return
	}
	i := quickAdjust(s, left, right)
	quickSort(s, left, i-1)
	quickSort(s, i+1, right)
}

func quickAdjust(s []int, left, right int) int {
	i, j := left, right

	// 每次以最左边元素为基准值的位置a,从右往左找第一个小于基准值的元素位置c,
	// 从左向右找第一个大于基准值的元素位置b,交换b和c,然后交换a和b
	// 保证位置b左边元素都小于b所在元素值,位置b右边元素都大于b所在元素值
	for i < j {
		for i < j && s[j] >= s[left] {
			j--
		}
		for i < j && s[i] <= s[left] {
			i++
		}
		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
	s[i], s[left] = s[left], s[i]
	return i
}
