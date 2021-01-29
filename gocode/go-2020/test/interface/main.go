package main

import (
	"fmt"
	"path/filepath"
)

type worker interface {
	work()
}

type person struct {
	name string
	worker
}

//func (p person) work() {
//	fmt.Println(p.name)
//}

func main() {
	var w worker = person{}
	fmt.Println(w)
	//w.work()
	var path = "/opt/dana/datax/errorData//"
	fmt.Println(filepath.Clean(path))
	var in interface{}
	if len(path) > 10 {
		in = len(path)
	}
	if in != nil {
		fmt.Println(fmt.Sprintf("%v", in))
	}
}
