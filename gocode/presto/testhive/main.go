package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	// "strings"
	"time"

	"github.com/uxff/gohive"
)

func main() {
	timeStr := time.Now().Format("2006-01-02")
	now, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	// 获取昨天零点的日期
	timeNow := now.AddDate(0, 0, -1).Unix()
	fmt.Println("time:", timeNow)
	fmt.Println("1-----------1")
	// 70.13 70.10 80.29 90.144 80.20 1.2
	db, err := gohive.Connect("192.168.80.29:10000", gohive.DefaultOptions)
	if err != nil {
		fmt.Println("hive连接错误：", err.Error())
		return
	}
	fmt.Println("2----------------2")
	defer db.Close()
	// sqltxt := "show databases"
	// sqltxt := "set hive.execution.engine"
	// dwddb.dwd_sc_1
	sqltxt := "show tables in odsdb"
	tt := time.Now()
	// sqltxt := "select count(1) from basedb.base_aitest  where nationality=53"
	// sqltxt := "show partitions odsdb.bee123"
	// sqltxt := "select count(*) from odsdb.bee123"
	version, err := db.SimpleQuery(sqltxt)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(version)
	fmt.Println("耗时：", time.Since(tt))
	// sqltxt = "analyze table odsdb.bee123 PARTITION(pt='2020-04-02 16:01:55.527') COMPUTE STATISTICS"
	// _, err = db.Exec(sqltxt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	tables := []string{}
	for _, v := range version {
		tables = append(tables, fmt.Sprintf("%v", v["tab_name"]))
	}
	for _, v := range tables {
		tt := time.Now()
		// sqltxt = fmt.Sprintf("select * from `odsdb`.`%s` limit 1", v)
		sqltxt = fmt.Sprintf("show create table `odsdb`.`%s`", v)
		version, err = db.SimpleQuery(sqltxt)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s : %v\n", v, version)
		if len(version) == 0 {
			fmt.Printf("表%s为空表\n", v)
		}
		for _, v := range version {
			loc := fmt.Sprintf("%v", v["createtab_stmt"])
			if strings.Contains(loc, "hdfs://") {
				fmt.Println(v["createtab_stmt"])
				break
			}
		}
		fmt.Println("耗时：", v, time.Since(tt))
		break
	}
	// fmt.Println(version)
	// 70.10 ycc.ods_oraclecq05
	// 80.20 chb.student
	// 1.5  odsdb.teacher
	// 70.13 odsdb.xtest
	// sqltxt := "desc formatted odsdb.xtest"
	// result, err := db.SimpleQuery(sqltxt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, v := range result {
	// 	datatype := fmt.Sprintf("%v", v["data_type"])
	// 	if strings.Contains(datatype, "lastDdlTime") {
	// 		comment := fmt.Sprintf("%v", v["comment"])
	// 		comment = strings.TrimSpace(comment)
	// 		fmt.Printf("%#v\n", comment)
	// 		break
	// 	}
	// }
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
