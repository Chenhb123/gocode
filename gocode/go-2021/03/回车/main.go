package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `123
456
789`
	s = strings.Replace(s, "\r", " ", -1)
	s = strings.Replace(s, "\n", " ", -1)
	fmt.Println(s)
}
