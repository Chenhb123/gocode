package main

func main() {
	x, y := 1, 2
	defer func(a int) {
		println("x,y的值分别为：", a, y)
	}(x)
	x += 100
	y += 100
	println(x, y)
	// 延迟调用可修改当前函数命名返回值，但其自身返回值会被抛弃
	println(test())
}

func test() (z int) {
	defer func() {
		println(z)
		z += 100
	}()
	return 100
}
