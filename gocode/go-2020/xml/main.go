package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// SConfig core-site.xml文件结构
type SConfig struct {
	XMLName  xml.Name    `xml:"configuration"` // 指定最外层的标签为configuration
	Property []SProperty `xml:"property"`      // 读取proeprty配置项

}

// SProperty property配置项
type SProperty struct {
	Name  string `xml:"name"`  // key
	Value string `xml:"value"` // value
}

func main() {
	file, err := os.Open("D:/share/vmdir/gopath/src/gocode/go-2020/xml/core-site.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	v := SConfig{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}
	for _, pro := range v.Property {
		if strings.Contains(pro.Name, "hadoop.rpc.protection") {
			fmt.Println("value:", pro.Value)
			break
		}
	}
}
