package main

import "fmt"

func main() {
	fmt.Println(increase(1))
}

func increase(d int) (ret int) {
	defer func() {
		d++
	}()

	return d
}
