package main

func main() {
	add := func(a, b int) int {
		return a + b
	}
	println(add(1, 2))
	testChannel()
}

func testChannel() {
	c := make(chan func(int, int) int, 2)
	c <- func(a, b int) int {
		return a + b
	}
	println((<-c)(1, 2))
}
