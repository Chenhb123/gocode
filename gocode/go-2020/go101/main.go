package main

// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

func main() {
	const s = "Go101.org"
	var a byte = 1 << len(s) / 128
	var b byte = 1 << len(s[:]) / 128
	println(a, b) // 4 0
}
