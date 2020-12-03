package main

import "log"

func main() {
	test(5, 0)
	defer catch()
	defer log.Println(recover())
	defer recover()
	panic("i am dead")
}

func catch() {
	log.Println("catch: ", recover())
}

func test(x, y int) {
	z := 0
	// recover必须在延迟函数中执行才能正常工作
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
	}()
	println(x, "/", y, "=", z)
}
