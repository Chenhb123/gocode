/*

Go提供的time包里面，按照下面几种场景做了区分，并分别提供一些API支持。主要场景如下所示：

1.超时一次之后，就不再使用的定时器，time.After()。
2.每隔一段时间，就需要使用一次的定时器，time.Tick()。
3.阻塞住goroutinue的定时器，time.Sleep()，准确来说这个不算一个定时器
4.可以自由控制定时器启动和关闭的定时器，time.Ticker()。

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	tchan := time.After(time.Second * 3)
	fmt.Println(time.Now().String(), "tchan type:%T", tchan)
	fmt.Println(time.Now().String(), "mark 1")
	fmt.Println(time.Now().String(), "tchan=", <-tchan) // channel里取出数据之后，会发现超时间是3秒
	fmt.Println(time.Now().String(), "mark 2")

	a := map[string]interface{}{
		"sss": 1,
		"bbb": 2,
	}
	fmt.Println(a["sss"])
}
