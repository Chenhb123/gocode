package main

import (
	"datatom/gin.v1"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"log"
)

func main() {
	var body interface{}
	body = gin.H{
		"query": gin.H{
			"term": gin.H{
				"id": "123456",
			},
		},
	}
	bodyByte, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	//bodyStr, _ = sjson.Set(bodyStr, "size", 1)
	js, err := simplejson.NewJson(bodyByte)
	if err != nil {
		fmt.Println("1111")
		log.Fatal(err)
	}
	js.SetPath([]string{"size"}, 1)
	encode, _ := js.Encode()
	bodyStr := string(encode)
	fmt.Println(bodyStr)

}
