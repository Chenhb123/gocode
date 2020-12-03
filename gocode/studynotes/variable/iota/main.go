package main

import "fmt"

const (
	c, cc = iota, 10 * iota
	d, dd
	e, ee
)

const (
	_ = 1 << (10 * iota)
	// KB .
	KB
	// MB .
	MB
	// GB .
	GB
)

func main() {
	fmt.Println(KB, MB, GB)
	fmt.Printf("%T,%v\n", d, d)
	fmt.Printf("%T,%v\n", ee, ee)
}
