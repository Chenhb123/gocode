package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("123")

	// var m = []int{1, 2, 3}
	// for i := range m {
	// 	go func() {
	// 		fmt.Print(i)
	// 	}()
	// }
	// //block main 1ms to wait goroutine finished
	// time.Sleep(time.Millisecond)

	// var m = map[int]int{1: 1, 2: 2, 3: 3}
	// for i := range m {
	// 	m[4] = 4
	// 	fmt.Printf("%d:%d ", i, m[i])
	// }

	// var m = map[int]int{1: 1, 2: 2, 3: 3}
	m := []int{1, 2, 3}
	//only del key once, and not del the current iteration key
	var o sync.Once
	for i, v := range m {
		if i < 2 {
			fmt.Printf("%d:%d ", i, v)
		}
		o.Do(func() {
			m = m[:2]
			// for _, key := range []int{1, 2, 3} {
			// 	if key != i {
			// 		fmt.Printf("when iteration key %d, del key %d\n", i, key)
			// 		// delete(m, key)
			// 		m = m[1:]
			// 		break
			// 	}
			// }
		})
	}

	// arr := [2]int{1, 2}
	// res := []*int{}
	// for _, v := range arr {
	// 	res = append(res, &v)
	// }
	// //expect: 1 2
	// fmt.Println(*res[0], *res[1])
	// //but output: 2 2
}
