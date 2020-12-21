package main

import "fmt"

/*

三个数组成4个数，这三个数只能在280中选，且必须包含280
比如：2880合法，而2220、3280不合法
*/

func main() {
	// 使用二叉树求解
	// 最上层为千位，从2、8中任选
	s := []int{2, 8, 0}
	var slist [][]int
	var temp, used []int
	circle(s, used, temp, slist)
	for _, v := range slist {
		str := ""
		for _, s := range v {
			str += fmt.Sprintf("%d", s)
		}
		fmt.Println(str)
	}
}

func circle(s, used, temp []int, slist [][]int) {
	if len(used) == 3 {
		if len(temp) < 4 {
			for _, v := range s {
				t := temp
				t = append(t, v)
				slist = append(slist, t)
			}
			return
		}
		slist = append(slist, temp)
		return
	}
	// 拼接
	for _, v := range s {
		if len(temp) == 0 && v == 0 {
			continue
		}
		target := false
		for _, u := range used {
			if u == v {
				target = true
				break
			}
		}
		if !target {
			used = append(used, v)
		}
		temp = append(temp, v)
	}
}
