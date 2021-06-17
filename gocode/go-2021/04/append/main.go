package main

import "fmt"

func main() {

	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))

	//s := []int{5}
	//s = append(s, 7)
	//s = append(s, 9)
	//x := append(s, 11)
	//y := append(s, 12, 12)
	//fmt.Println(s, x, y)

	//s := make([]int, 0)
	//
	//oldCap := cap(s)
	//
	//for i := 0; i < 2048; i++ {
	//	s = append(s, i)
	//
	//	newCap := cap(s)
	//
	//	if newCap != oldCap {
	//		fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n",
	//			0, i-1, oldCap, i, newCap)
	//		oldCap = newCap
	//	}
	//}
	//str := fmt.Sprintf("zip")
	//for _, v := range s {
	//	str = fmt.Sprintf("%s %v", str, v)
	//}
	//fmt.Println(str)
}

//// go 1.14.6 src/runtime/slice.go:76
//// et:slice元素类型,old:旧的slice,cap:所需的新最小容量
//func growslice(et *_type, old slice, cap int) slice {
//	// ……
//	newcap := old.cap
//	doublecap := newcap + newcap
//	if cap > doublecap {
//		newcap = cap
//	} else {
//		// ……
//	}
//	// ……
//	// Specialize for common values of et.size.
//	// For 1 we don't need any division/multiplication.
//	// For sys.PtrSize, compiler will optimize division/multiplication into a shift by a constant.
//	// For powers of 2, use a variable shift.
//	switch {
//	case et.size == 1:
//		// ……
//	case et.size == sys.PtrSize:
//		lenmem = uintptr(old.len) * sys.PtrSize
//		newlenmem = uintptr(cap) * sys.PtrSize
//		capmem = roundupsize(uintptr(newcap) * sys.PtrSize)
//		overflow = uintptr(newcap) > maxAlloc/sys.PtrSize
//		newcap = int(capmem / sys.PtrSize)
//	case isPowerOfTwo(et.size):
//		// ……
//	default:
//		// ……
//	}
//}
