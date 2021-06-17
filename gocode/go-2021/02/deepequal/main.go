package main

import (
	"fmt"
)

type S struct {
	a, b, c string
}

func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})
	a, b := &S{"a", "b", "c"}, &S{"a", "b", "c"}
	//fmt.Println(reflect.DeepEqual(x, y))
	fmt.Println(x == y)
	fmt.Println(*a == *b)
}
