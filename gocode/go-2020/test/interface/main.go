package main

import (
	"fmt"
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
	w.work()
}
