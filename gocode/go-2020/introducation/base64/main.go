package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

/*
base64编解码
*/

func main() {
	// 需要处理的字符串
	message := "i konw you meaning but i can not allow you do what you want to do"
	// 编码消息
	encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))
	// 输出编码完成的消息
	fmt.Println(encodeMessage)
	// 解码消息
	data, err := base64.StdEncoding.DecodeString(encodeMessage)
	if err != nil {
		log.Fatal(err)
	}
	// 打印解码完成的数据
	fmt.Println(string(data))
}
