package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	list := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	lFields := strings.Fields(string(line))
	for _, l := range lFields {
		num, _ := strconv.Atoi(l)
		list = append(list, num)
	}
	//fmt.Println(list)
	length := len(list)
	for i := length - 1; i >= 1; {
		if list[i] == -1 {
			continue
		}
		var temp int
		for j := i - 1; j >= 0; j-- {
			if list[j] == -1 {
				continue
			}
			temp += list[j]
			if temp > list[i] {
				i--
				break
			} else if temp == list[i] {
				list = append(list, list[i-1]*2)
				for t := i - 1; t >= j; t-- {
					list[t] = -1
				}
				i--
				break
			}
			if j == 0 {
				i--
			}
		}
	}
	for _, v := range list {
		if v != -1 {
			fmt.Printf("%d ", v)
		}
	}
}
