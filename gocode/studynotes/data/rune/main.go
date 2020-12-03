package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	s := "喵喵喵喵"
	fmt.Println("len(s)=", utf8.RuneCountInString(s))
	p := reflect.ValueOf(&s)
	v := p.Elem()
	fmt.Printf("%v\n", v.String())
	// transform("s: %x\n", s)
}

func transform(format string, ptr interface{}) {

	// fmt.Printf(format, h)
}
