package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "阳阳"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: [%c]\n", i, s[i])
	}
	for i, v := range s {
		fmt.Printf("%d: [%c]\n", i, v)
	}
	// test()
	// s := "杨过\x61\142\u0041"
	// fmt.Printf("%s\n", s)
	// fmt.Printf("% x, len(s)= %d\n", s, len(s))
	// s = "ab" +
	// 	"cd"
	// println(s == "abcd")
	// println(s > "abc")
}

func test() {
	s := "abcdefgh"
	s1 := s[:3]
	s2 := s[1:4]
	s3 := s[2:]
	println(&s)
	println(&s1)
	println(s1, s2, s3)
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
}
