package main

import "fmt"

// 同构字符串

func main() {
	s1 := []rune("abc")
	s2 := []rune("cdf")
	m := make(map[rune]rune)
	for i := 0; i < len(s1); i++ {
		if value, ok := m[s1[i]]; !ok {
			// 不存在
			m[s1[i]] = s2[i]
		} else {
			if value != s2[i] {
				return
			}
		}
	}
	fmt.Println("2-------------2")
}
