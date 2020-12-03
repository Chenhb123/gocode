package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"time"
)

var Number int
var Result string

func main() {
	config := hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  8,
		RequestVolumeThreshold: 5,
		SleepWindow:            1,
		ErrorPercentThreshold:  30,
	}
	hystrix.ConfigureCommand("test", config)
	cbs, _, _ := hystrix.GetCircuit("test")
	defer hystrix.Flush()
	for i := 0; i < 15; i++ {
		start := time.Now()
		Number = i
		hystrix.Do("test", run, getFallBack)
		fmt.Println("请求次数：", i+1, "，用时：", time.Now().Sub(start), "，Result：", Result,
			"，熔断器开启状态：", cbs.IsOpen(), "，请求是否允许：", cbs.AllowRequest())
	}
	time.Sleep(20 * time.Second)
}

func run() error {
	Result = "RUNNING"
	if Number > 10 {
		return nil
	}
	time.Sleep(2000 * time.Millisecond)
	if Number%2 == 0 {
		return nil
	} else {
		return fmt.Errorf("请求失败")
	}
}

func getFallBack(err error) error {
	Result = "FALLBACK"
	return nil
}
