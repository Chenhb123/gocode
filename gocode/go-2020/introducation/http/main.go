package main

import "net/http"

/*
实现一个最简单的http服务
*/

func main() {
	http.Handle("/", http.FileServer(http.Dir("D:/share/vmdir/gopath/src/github.com")))
	http.ListenAndServe(":8080", nil)
}
