package main

import "log"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		} else {
			log.Fatalln("fatal")
		}
	}()
	defer func() {
		panic("i can replace")
	}()
	panic("i am dead")
	// println("exit") // 不会被执行
}
