package main

import (
	"fmt"
	"regexp"
)

func main() {
	email := "1219062061@163.com"
	target, _ := regexp.MatchString("^\\w+([-+.]\\w+)*@(163|qq|datatom|gmail)\\.com$", email)
	if !target {
		// 邮箱不合法
		fmt.Printf("邮箱非法:%s\n", email)
	}
	fmt.Printf("合法:%s\n", email)
}
