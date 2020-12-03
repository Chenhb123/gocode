package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/uxff/gohive"
)

func main() {
	var ip, port, sql string
	flag.StringVar(&ip, "i", "127.0.0.1", "IP地址,默认为127.0.0.1")
	flag.StringVar(&port, "p", "10000", "端口号,默认为10000")
	flag.StringVar(&sql, "e", "show databases", "sql语句,默认为show databases")
	flag.Parse()
	fmt.Println("ip:", ip)
	fmt.Println("port:", port)
	fmt.Println("sql:", sql)
	timeStr := time.Now().Format("2006-01-02")
	tt := time.Now()
	fmt.Println("当前时间:", timeStr)
	str := fmt.Sprintf("%s:%s", ip, port)
	db, err := gohive.Connect(str, gohive.DefaultOptions)
	if err != nil {
		fmt.Println("hive连接错误：", err.Error())
		return
	}
	defer db.Close()
	version, err := db.SimpleQuery(sql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
	fmt.Println("耗时：", time.Since(tt))
}
