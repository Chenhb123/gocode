package main

import "github.com/skip2/go-qrcode"

func main() {
	qrcode.WriteFile("https://juejin.im/post/5dd292ef518825638b753975", qrcode.Medium, 256, "blog_qrcode.png")
}
