package main

import "fmt"

func main() {
	size := 786432
	fmt.Println(fmt.Sprintf("%.2f KB", float64(size)/1024))
}
