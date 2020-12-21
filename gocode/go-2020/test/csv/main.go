package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	mapStr := []map[string]interface{}{
		{"bigint": 1, "bit": true, "bool": 1, "boolean": 1, "varchar_1": "香港"},
		{"bigint": 1, "bit": true, "bool": 1, "boolean": 1, "varchar_1": "香港"},
		{"bigint": 1, "bit": true, "bool": 1, "boolean": 1, "varchar_1": "香港"},
		{"bigint": 1, "bit": true, "bool": 1, "boolean": 1, "varchar_1": "香港"},
		{"bigint": 1, "bit": true, "bool": 1, "boolean": 1, "varchar_1": "香港"},
	}
	var data [][]string
	if len(mapStr) <= 0 {
		fmt.Println("数据为空")
		return
	}
	var header []string
	for i := range mapStr[0] {
		header = append(header, i)
	}
	data = append(data, header)
	for _, m := range mapStr {
		var line []string
		for _, v := range header {
			line = append(line, fmt.Sprintf("%v", m[v]))
		}
		data = append(data, line)
	}

	csvFile, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, v := range data {
		writer.Write(v)
	}
	writer.Flush()
}
