package main

func main() {
	f1, f2 := test123(123)
	f1()
	f2()
	// f := test(0x100)
	// f()
	// res := testX()
	// for _, v := range res {
	// 	v()
	// }
}

func test(x int) func() {
	println(&x)
	return func() {
		println(&x, x)
	}
}

func testX() []func() {
	res := []func(){}
	for i := 0; i < 2; i++ {
		x := i
		res = append(res, func() {
			println(x)
		})
	}
	return res
}

func test123(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}
