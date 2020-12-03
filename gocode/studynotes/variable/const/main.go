package main

import (
	"fmt"
	"unsafe"
)

const (
	x int = 10
	y
	a = "hello"
	b
)

func main() {
	fmt.Println(x, y)
	fmt.Println(a, b)
	ptrSize := unsafe.Sizeof(uintptr(0))
	strSize := len("hello")
	fmt.Println(ptrSize, strSize)
}
