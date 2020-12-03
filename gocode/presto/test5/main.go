package main

import (
	"fmt"
	"log"
	"strings"

	tools "datatom.com/tools/httpdo"
)

func main() {
	path := "/root/TDH/TDH-Client11"
	command := fmt.Sprintf("if [ -d %s ]; then\n echo \"exist\"\n else echo \"not exist\"\n fi", path)
	res, err := tools.ExecCmd("192.168.90.21", 22, "root", "datatom", command)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("res:", res)
	if strings.Contains(res, "not") {
		log.Fatal(fmt.Errorf("客户端路径不正确：version:%s, path:%s", "3.2.1", path))
	}
	fmt.Println("end")
}
