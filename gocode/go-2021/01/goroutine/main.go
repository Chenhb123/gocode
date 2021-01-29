package main

import (
	"fmt"
	"time"
)

/*
假设有4个 goroutine，编号为1，2，3，4。每秒钟会有一个 goroutine 打印出它自己的编号。
现在让你写一个程序，要求输出的编号总是按照1，2，3，4这样的顺序打印。类似下图，
      --------------3---------------
      |                             |
      |                             |
      |                             |
      |                             |
      4                             2
      |                             |
      |                             |
      |                             |
      |                             |
      --------------1----------------

                     ^
                     |
                     |起始点
                     |
*/

type token struct{}

func main() {
	var num int
	num = 4
	var chs []chan token
	for i := 0; i < num; i++ {
		chs = append(chs, make(chan token))
	}
	// 同时起4个goutine
	/*
		worker(0, chs[0], chs[1])
		worker(1, chs[1], chs[2])
		worker(2, chs[2], chs[3])
		worker(3, chs[3], chs[0])
	*/
	for i := 0; i < num; i++ {
		go worker(i, chs[i], chs[(i+1)%num])
	}
	// 给起始点通道赋值
	chs[0] <- struct{}{}
	select {}
}

func worker(id int, ch, next chan token) {
	for {
		token := <-ch
		fmt.Println(id + 1)
		time.Sleep(time.Second)
		next <- token
	}
}
