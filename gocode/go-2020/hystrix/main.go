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
		Timeout:                1000, // 超时时间设置  单位毫秒
		MaxConcurrentRequests:  8,    // 最大请求数
		SleepWindow:            1,    // 过多长时间，熔断器再次检测是否开启,单位毫秒
		RequestVolumeThreshold: 5,    // 请求阈值  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
		ErrorPercentThreshold:  30,   // 错误率 %
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
	if Number > 5 {
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
	//fmt.Println("err:", err.Error())
	return nil
}
