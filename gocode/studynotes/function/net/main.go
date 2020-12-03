package main

import (
	"fmt"
)

func main() {
	res := test()
	for i := 0; i < len(res); i++ {
		fmt.Println(&res[i])
	}
	// addrs, err := net.InterfaceAddrs()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, value := range addrs {
	// 	if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	// 		if ipnet.IP.To4() != nil {
	// 			fmt.Println(ipnet.IP.String())
	// 		}
	// 	}
	// }
}

func test() []int {
	var res []int
	for i := 0; i < 3; i++ {
		res = append(res, i)
	}
	for i := 0; i < len(res); i++ {
		fmt.Println(&res[i])
	}
	return res
}
