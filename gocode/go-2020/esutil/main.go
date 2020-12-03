package main

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	var id = "123456"
	ibytes, _ := json.Marshal(id)
	array := gjson.ParseBytes(ibytes).Array()
	for _, v := range array {
		fmt.Println(v.String())
	}
}
