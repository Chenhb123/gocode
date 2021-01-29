package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = `,./[]\<>?:"{}|+_)(*&^%$#@!，。、；‘【】、=-·《》？：“{}|——+）（&*%……#￥！@~"`
	if IsTableSpe(str) {
		fmt.Println(true)
	}
}

func IsTableSpe(describe string) bool {
	str := ";" + "'" + "--"

	if strings.ContainsAny(describe, str) {
		return true
	}

	return false
}
