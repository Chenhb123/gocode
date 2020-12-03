package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

// Poc poc
type Poc struct {
	Method     string  `json:"method"`
	Headers    Headers `json:"headers"`
	Path       string  `json:"path"`
	Body       string  `json:"body"`
	Expression int     `json:"expression"`
}

// Headers Headers
type Headers struct {
	UserAgent      string `json:"useragent"`
	Accept         string `json:"acept"`
	XForwardedFor  string `json:"xforwardedfor"`
	ContentType    string `json:"contenttype"`
	Referer        string `json:"referer"`
	AcceptLanguage string `json:"acceptlanguage"`
	Cookie         string `json:"cookie"`
}

func main() {
	//这里接收用户命令行中的-m参数
	var module string
	flag.StringVar(&module, "m", "H1", "module: all")
	flag.Parse()
	//把用户命令行中的-m的值传给getpocinfo函数
	//但是getpocinfo函数接收的是Poc类型
	// fmt.Println(module)
	var poc Poc
	err := json.Unmarshal([]byte(module), &poc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", poc)
	a, b, c, d := getpocinfo(poc)
	fmt.Println(a + ":" + b + ":" + c + ":" + d)

}
func getpocinfo(id Poc) (method string, path string, header string, body string) {
	method = id.Method
	path = id.Path
	header = id.Headers.UserAgent
	body = id.Body
	return method, path, header, body
}

// H1 H1
var H1 = Poc{Method: "GET", Path: "/resin-doc/viewfile/?file=index.jsp", Body: "", Expression: 200, Headers: Headers{UserAgent: "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"}}
