package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/prestodb/presto-go-client/presto"
)

func main() {
	s := "http://root@192.168.2.155:9999?catalog=hive&schema=chb_test"

	db, err := sql.Open("presto", s)
	if err != nil {
		fmt.Println("33")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("连接失败")
	}

	sqltxt := "show create table prestotest.comtb"
	rows, err := db.Query(sqltxt)
	if err != nil {
		fmt.Println("query err")
	}
	var res string
	for rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			fmt.Println("scan err")
			fmt.Println(err.Error())
		}
	}
	pre := strings.Index(res, "(")
	res = res[pre+1:]
	pre = strings.Index(res, ")")
	res = res[:pre]
	res = strings.Trim(res, "\n")
	resArr := strings.Split(res, "\n")
	for _, v := range resArr {
		v = strings.TrimSpace(v)
		vArr := strings.Split(v, " ")
		fmt.Println(vArr)
	}
	// fmt.Printf("%#v\n", res)
	// result, err := sqlrows2Maps(rows)
	// if err != nil {
	// 	fmt.Println("22")
	// }
	// fmt.Println(rows)
	// sqltxt := "show tables"
	// //fmt.Println(sqltxt)
	// res, err := db.Query(sqltxt)
	// fmt.Println(res)
	// return
}
