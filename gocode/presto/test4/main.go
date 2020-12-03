package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	name := tableRename("中国$#@chb122")
	fmt.Println(name)
}

func rRotate(nums []int, k int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i+k < len(nums) {
			res[i+k] = nums[i]
			continue
		}
		len := i + k - len(nums)
		res[len] = nums[i]
	}
	return res
}

func tableRename(tableName string) string {
	tableRune := []rune(tableName)
	// 表名是否包含特殊字符，若包含则替换为字符索引+1
	speRune := []rune{'~', '`', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '=', '+', '.'}
	for i := range tableRune {
		for _, spe := range speRune {
			if spe == tableRune[i] {
				tableRune[i] = '_'
			}
		}
		// 是否包含中文
		if unicode.Is(unicode.Han, tableRune[i]) {
			now := time.Now().Format("2006_01_02_15_04_05")
			tableName = "default" + now
			return tableName
		}
	}
	tableName = string(tableRune)
	fmt.Println("hive表名：", tableName)
	return tableName
}
