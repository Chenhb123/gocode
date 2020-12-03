package main

import (
	"fmt"

	valid "gopkg.in/asaskevich/govalidator.v4"
)

func main() {
	str := "Hello World!"
	str = valid.Reverse(str)
	fmt.Println(str)
	str = ` Today is 6nd,May,2019, everyone
	every every day feels good but but but
	the summer is becoming becoming, the day
	will be hotter hotter, water needs more more
	more, timer should be more more more.`
	count := valid.GetLines(str)
	fmt.Println(len(count))
	str = "A B C D E F"
	str = valid.Truncate(str, 3, "!")
	fmt.Println(str)
}
