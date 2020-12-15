package main

import (
	"fmt"
	"gocode-chb/gocode/binarytree/another"
)

var test = map[string]int{
	"1": 1,
}

func main() {
	dir.Test()
	var x int
	inc := func() int {
		x++
		return x
	}
	fmt.Println(func() (a, b int) {
		return inc(), inc()
	}())
	var 测试 = "123"
	fmt.Println(测试)
	//dir.天气()
}
