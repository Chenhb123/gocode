package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/prestodb/presto-go-client/presto"
)

// 数据源配置
var dsn = "http://root@192.168.80.29:9999?catalog=hive&schema=wt"

func main() {
	db, err := open()
	if err != nil {
		log.Fatal(err)
	}
	defer close(db)
	sqltxt := "select * from xly_test"
	tables, err := query(db, sqltxt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tables)
}

// 查询
func query(db *sql.DB, sqltxt string) ([]map[string]interface{}, error) {
	var tables []map[string]interface{}
	rows, err := db.Query(sqltxt)
	if err != nil {
		return tables, err
	}
	defer rows.Close()
	tables, err = Sqlrows2Maps(rows)
	if err != nil {
		return tables, err
	}
	return tables, nil
}

// 打开连接
func open() (*sql.DB, error) {
	db, err := sql.Open("presto", dsn)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}

// 关闭连接
func close(db *sql.DB) {
	db.Close()
}

// Sqlrows2Maps sql查询结果rows转为maps
func Sqlrows2Maps(rws *sql.Rows) ([]map[string]interface{}, error) {
	rowMaps := make([]map[string]interface{}, 0)
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
