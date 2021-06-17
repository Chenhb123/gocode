package main

import "fmt"

// 简单选择排序

func main() {
	s := []int{1, 2, 3, 2, 2, 3, 4, 3, 3}
	selectSort(s)
	fmt.Println(s)
}

/*
稳定：如果a原本在b前面，而a=b，排序之后a仍然在b的前面;
不稳定：如果a原本在b的前面，而a=b，排序之后 a 可能会出现在 b 的后面
*/
func selectSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[min] > s[j] {
				min = j
			}
		}
		if min != i {
			s[min], s[i] = s[i], s[min]
		}
	}
}
