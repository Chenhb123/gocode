package main

import "fmt"

type S struct {
	name string
}

func main() {
	mm := map[string]S{"y": S{}}
	c, ok := mm["y"]
	if ok {
		fmt.Println(c.name)
	}
	ss := []S{{"one"}, {"two"}, {"three"}}
	ss[0].name = "copy"

}
