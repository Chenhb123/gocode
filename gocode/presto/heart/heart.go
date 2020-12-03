package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func unquote(s string) string {
	s, err := strconv.Unquote(`"` + s + `"`)
	if err != nil {
		panic(err)
	}
	return s
}

// Print text heart
// Author: ShixiangWang
// LICENSE: MIT
// Reference: https://blog.csdn.net/su_bao/article/details/80355001
func main() {
	// MYWORD My word
	var head string
	var tail string
	var MYWORD string
	var sep string
	var zoom float64
	flag.StringVar(&head, "head", "There are some words I wana tell you:", "A sentence printed on the head")
	flag.StringVar(&tail, "tail", "\t\t\t\t--- Your lover", "A sentence printed on the tail")
	flag.StringVar(&MYWORD, "words", "Dear, I love you forever!", "The words you want to talk")
	flag.StringVar(&sep, "sep", " ", "The separator")
	flag.Float64Var(&zoom, "zoom", 1.0, "Zoom setting")
	flag.Parse()

	//fmt.Printf("Words: %T\n", MYWORD)
	//fmt.Printf("tail: %T\n", tail)

	chars := strings.Split(MYWORD, sep)

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(unquote(head))
	fmt.Println()
	time.Sleep(time.Duration(1) * time.Second)

	//fmt.Printf("chars: %T\n", chars)       // []string
	//fmt.Printf("chars[0]: %T\n", chars[0]) // string

	for _, ch := range chars {

		//fmt.Printf("ch[0]: %T\n", ch[0]) // uint8, i.e. byte

		char := []rune(ch)

		//fmt.Println(char)
		//fmt.Printf("char[0]: %T\n", char[0]) // int32, i.e. rune

		allChar := make([]string, 0)
		for y := 12 * zoom; y > -12*zoom; y-- {
			lst := make([]string, 0)
			lstCon := ""
			for x := -30 * zoom; x < 30*zoom; x++ {
				x2 := float64(x)
				y2 := float64(y)
				formula := math.Pow(math.Pow(x2*0.04/zoom, 2)+math.Pow(y2*0.1/zoom, 2)-1, 3) - math.Pow(x2*0.04/zoom, 2)*math.Pow(y2*0.1/zoom, 3)
				if formula <= 0 {
					index := int(x) % len(char)
					if index >= 0 {
						lstCon += string(char[index])
					} else {
						lstCon += string(char[int(float64(len(char))-math.Abs(float64(index)))])
					}

				} else {
					lstCon += " "
				}
			}
			lst = append(lst, lstCon)
			allChar = append(allChar, lst...)
		}

		for _, text := range allChar {
			fmt.Printf("%s\n", text)
		}
	}
	time.Sleep(time.Duration(1) * time.Second)

	fmt.Println(unquote(tail))
	time.Sleep(1 * time.Minute)
}
