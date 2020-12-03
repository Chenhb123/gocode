package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now().Unix()
	timeStr := time.Now().Format("2006-01-02")
	t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	startTime := t2.AddDate(0, 0, 1).Unix()
	interval := startTime - timeNow
	fmt.Println(interval)
}
