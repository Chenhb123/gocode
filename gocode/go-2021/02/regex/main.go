package main

import (
	"fmt"
	"regexp"
)

/*
100
0
1
0
1
*/

func main() {
	s := `100
00
0
1
0

0
1`
	pattern := regexp.MustCompile("0\n+0+") //A
	s = pattern.ReplaceAllString(s, "0")
	fmt.Println(s)
}
