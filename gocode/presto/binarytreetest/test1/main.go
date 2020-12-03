package main

import (
	"encoding/json"
	"fmt"
	binarytree "gocode/binarytree/tree"
	"log"
	"os"

	simplejson "github.com/bitly/go-simplejson"
)

/*
		1
	2		2
   3          3
  4             4

*/

var content = `
{
	"job": {
	 "setting": {
	 "speed": {
		  "channel": 32
	 },
	 "errorLimit": {
		 "record": 0,
		 "percentage": 0.02
	 }
 },
	 "content": [{
		 "reader": {
			 "name": "txtfilereader",
	 "skipHeader":false,
			 "parameter": {
				 "path": ["/home/csv/testkt.csv"],
				 "encoding": "UTF-8",
				 "column": [{"index":0,"type":"string"},{"index":1,"type":"string"},{"index":2,"type":"string"}],
				 "fieldDelimiter":","
			 }
		 },
		 "writer": 
{
"name": "hdfswriter",
"parameter": {
 "defaultFS": "hdfs://192.168.2.100:8020",
 "fileType": "text",
 "path": "/zzzzz",
 "fileName": "demo",
 "column":[{"name":"dsource_id","type":"string"},{"name":"user_id","type":"string"},{"name":"dsource_name","type":"string"}],
 "writeMode": "append",
 "fieldDelimiter": "\t"
}
}

	 }]
 }
}

`

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

func testJSON(content string) (string, error) {
	path := ""
	js, err := simplejson.NewJson([]byte(content))
	if err != nil {
		return path, err
	}
	infos, err := js.Get("job").Get("content").Array()
	if err != nil {
		return path, err
	}
	bytes, err := json.Marshal(infos[0])
	if err != nil {
		return path, err
	}
	js, err = simplejson.NewJson(bytes)
	if err != nil {
		return path, err
	}
	path, err = js.Get("writer").Get("parameter").Get("path").String()
	if err != nil {
		return path, err
	}
	return path, nil
}

func test(res []int) *binarytree.BinaryTree {
	var root *binarytree.BinaryTree
	length := len(res)
	if length == 0 {
		return root
	}
	mid := length / 2
	root = binarytree.New(res[mid])
	if length == 1 {
		root.Left = nil
		root.Right = nil
	} else if length == 2 {
		root.Left = binarytree.New(res[0])
		root.Right = nil
	} else {
		lr := res[:mid]
		rr := res[mid+1:]
		root.Left = test(lr)
		root.Right = test(rr)
	}
	return root
}
