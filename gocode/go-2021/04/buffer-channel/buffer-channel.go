package main

import (
	"fmt"
	"sync"
)

func main() {
	//wg := sync.WaitGroup{}
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(wg sync.WaitGroup, i int) {
	//		fmt.Printf("i:%d", i)
	//		wg.Done()
	//	}(wg, i)
	//}
	//wg.Wait()
	//fmt.Println("exit")

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // len=10 cap=10
	s1 := slice[2:5]                             // [2, 3, 4] len=3 cap=8
	s2 := s1[2:6:7]                              // [4, 5, 6, 7] len=4 cap=5

	s2 = append(s2, 100) // [4, 5, 6, 7, 100]  slice[8]=100
	s2 = append(s2, 200) // [4, 5, 6, 7, 100, 200] len=5 cap=8

	s1[2] = 20 // slice[4]=20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
	m := sync.Map{}
	fmt.Println(m)

}
