package main

import (
	"fmt"
	"math"
)

// "reflect"
// "unsafe"

func countPrimes(n int) int {
	var sum int
	if n < 3 {
		return sum
	}
	isPrim := make([]bool, n)
	for i := range isPrim {
		isPrim[i] = true
	}
	isPrim[0], isPrim[1] = false, false
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if isPrim[i] {
			for j := i * i; j < n; j += i {
				isPrim[j] = false
			}
		}
	}
	for _, v := range isPrim {
		if v {
			sum++
		}
	}
	return sum
}

func main() {
	fmt.Println(countPrimes(10))
	// s := make([]int, 1)
	// fmt.Printf("dataPtr: %p len: %d, cap: %d %v\n", s, len(s), cap(s), s)
	// s = append(s, 1)
	// fmt.Printf("dataPtr: %p len: %d, cap: %d %v\n", s, len(s), cap(s), s)

	// // Delete 删除索引为i的元素
	// a := make([]map[int]int, 0)
	// for i := 0; i < 10; i++ {
	// 	m := make(map[int]int)
	// 	m[i] = i
	// 	a = append(a, m)
	// }
	// fmt.Println(a)
	// i := 6
	// if i < len(a)-1 {
	// 	copy(a[i:], a[i+1:])
	// }
	// a[len(a)-1] = nil // or the zero value of T
	// a = a[:len(a)-1]
	// fmt.Println(a)

	// // Cut 将索引i~j范围(不包括j)的元素去除
	// a := make([]map[int]int, 0)
	// for i := 0; i < 10; i++ {
	// 	m := make(map[int]int)
	// 	m[i] = i
	// 	a = append(a, m)
	// }
	// fmt.Println(a)
	// i, j := 6, 2
	// copy(a[i:], a[j:])
	// for k, n := len(a)-j+i, len(a); k < n; k++ {
	// 	a[k] = nil // or the zero value of T
	// }
	// a = a[:len(a)-j+i]
	// fmt.Println(a)

	// s := make([]int, 5)
	// x := s[:]
	// print(x)
	// arr := [5]int{}
	// y := arr[:]
	// print(y)

	// a := make([]int, 5)
	// s := a
	// invalid operation: s == a (slice can only be compared to nil)

	// print(a == s)

	// a := make([]int, 1, 3)
	// fmt.Printf("%p", a)
	// // reflect.SliceHeader 为 slice运行时数据结构
	// sh := (*reflect.SliceHeader)(unsafe.Pointer(&a))

	// fmt.Printf("slice header: %#v\naddress:\na:%p\n&a[0]:%p\n&a:%p\nsh:%p\nsh.Data:%p",
	// 	sh, a, &a[0], &a, sh, &sh.Data)

}
