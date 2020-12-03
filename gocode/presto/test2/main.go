package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

// DiskStatus .
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func main() {
	num := []int{4, 3, 2, 1}
	sort.Ints(num)
	fmt.Println(num)
	log.Fatal("测试一下")
	a, b := 1, 4
	fmt.Println(math.Abs(float64(a - b)))
	// res := make(chan int, 3)
	// var wg sync.WaitGroup

	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		defer wg.Done()
	// 		res <- n

	// 	}(i)
	// }
	// wg.Wait()
	// close(res)

	// for v := range res {
	// 	fmt.Println(v)
	// }
	// fmt.Println("main is end")

	// disk := DiskUsage("/")
	// fmt.Println(disk)
	// files := "/home/pkg/file.txt"
	// f := path.Base(files)
	// fmt.Println(f)
}

// disk usage of path/disk
// func DiskUsage(path string) (disk DiskStatus) {
// 	fs := syscall.Statfs_t{}
// 	err := syscall.Statfs(path, &fs)
// 	if err != nil {
// 		return
// 	}
// 	disk.All = fs.Blocks * uint64(fs.Bsize)
// 	disk.Free = fs.Bfree * uint64(fs.Bsize)
// 	disk.Used = disk.All - disk.Free
// 	return
// }
