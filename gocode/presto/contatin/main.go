package main

import (
	"fmt"
	"strings"

	"datatom.com/ants/httpdo/access"
)

func main() {
	file := "chb-mysql-hive.json"
	command := fmt.Sprintf("cat %s|grep 'defaultFS'", file)
	res, err := access.ExeCommand(command)
	if err != nil {
		fmt.Println("查询defaultFs出错：", err.Error())
		return
	}
	str := `hdfs://192.168.2.80`
	t := strings.Contains(res, str)
	fmt.Println(t)
}
