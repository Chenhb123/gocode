package main

import "runtime/debug"

func main() {
	defer func() {
		if recover() != nil {
			debug.PrintStack()
		}
	}()
	test()
}

func test() {
	panic("i am dead")
}
