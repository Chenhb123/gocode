package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/uxff/gohive"
)

func main() {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
	// now, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	// 获取昨天零点的日期
	// timeNow := now.AddDate(0, 0, -1).Unix()
	// fmt.Println("time:", timeNow)
	fmt.Println("1-----------1")
	// db, err := gohive.Connect("192.168.70.13:10000", gohive.DefaultOptions)
	db, err := gohive.Connect("192.168.90.144:10000", gohive.DefaultOptions)
	if err != nil {
		fmt.Println("hive连接错误：", err.Error())
		return
	}
	fmt.Println("2----------------2")
	defer db.Close()
	// sqltxt := "desc formatted odsdb.testkt"
	// sqltxt := "select * from odsdb.test "
	// sqltxt := "desc database `odsdb`"
	sqltxt := "desc formatted odsdb.ods_testcl"
	result, err := db.SimpleQuery(sqltxt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// for _, v := range result {
	// 	datatype := fmt.Sprintf("%v", v["data_type"])
	// 	if strings.Contains(datatype, "lastDdlTime") {
	// 		comment := fmt.Sprintf("%v", v["comment"])
	// 		comment = strings.TrimSpace(comment)
	// 		fmt.Printf("%#v\n", comment)
	// 		break
	// 	}
	// }
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// fmt.Println(result)
	// var dataType string
	// var flag bool
	// for _, v := range result {
	// 	for _, stmt := range v {
	// 		dataType = fmt.Sprintf("%v", stmt)
	// 		if strings.Contains(dataType, "hdfs://") {
	// 			fmt.Println(dataType)
	// 			flag = true
	// 			break
	// 		}
	// 	}
	// 	if flag {
	// 		break
	// 	}
	// 	// dataType := fmt.Sprintf("%v", v["createtab_stmt"])
	// 	// fmt.Println(dataType)
	// }
	// hdfsURL := strings.Replace(dataType, "'", "", -1)
	// // location:hdfs://dn100:8020/apps/hive/warehouse/odsdb.db/testmysql
	// index := strings.Index(hdfsURL, "hdfs://")
	// subres := hdfsURL[index+7:]
	// pre := strings.Index(subres, "/")
	// hdfsURL = subres[pre:]
	// command := fmt.Sprintf("hadoop fs -du -s %s", hdfsURL)
	// resbuf, err := httpdo.ExecCmd("192.168.80.20", 22, "root", "datatom", command)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// resArr := strings.Split(resbuf, "\n")
	// var subArr string
	// for _, v := range resArr {
	// 	if strings.Contains(v, hdfsURL) {
	// 		subArr = v
	// 	}
	// }
	// resArr = strings.Split(subArr, " ")
	// totalSize := resArr[0]
	// totalSize = strings.TrimSpace(totalSize)
	// if totalSize == "" || totalSize == "0" {
	// 	log.Fatal(err)
	// }
	// num, err := strconv.ParseUint(totalSize, 10, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(num)
	// ch <- sqltxt
}

//sqlrows2Maps sql查询结果rows转为maps
func sqlrows2Maps(rws *sql.Rows) ([]map[string]interface{}, error) {

	var rowMaps []map[string]interface{}

	var columns []string
	columns, err := rws.Columns()
	if err != nil {
		return rowMaps, err
	}

	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))
	for i := range values {
		scans[i] = &values[i]
	}

	for rws.Next() {
		_ = rws.Scan(scans...)
		each := map[string]interface{}{}
		for i, col := range values {
			each[columns[i]] = string(col)
		}

		rowMaps = append(rowMaps, each)
	}

	return rowMaps, nil
}
