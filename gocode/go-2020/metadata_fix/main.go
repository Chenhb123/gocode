package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"datatom.com/metadata/common"
	"datatom.com/metadata/httpdo/esutil"
	metasource "datatom.com/metadata/source"
	"datatom/gin.v1"
)

func main() {
	fixLayer()
}

// 4.6->4.6.1分层数据修正
/*
meta_layer -> meta_sourcedb -> meta_sourcetb


*/
func fixLayer() error {
	// 获取分层数据
	fmt.Println("开始获取分层数据:")
	var layers []metasource.MetaLayer
	bytes, err := esutil.SearchAll(common.DataBaseMetaData, common.TbMetaLayer)
	if err != nil {
		fmt.Println("获取分层数据失败:")
		log.Fatal(err)
	}
	data := esutil.GetSource(bytes)
	for _, v := range data {
		var lay metasource.MetaLayer
		err = json.Unmarshal([]byte(v), &lay)
		if err != nil {
			fmt.Println("解析分层数据失败")
			log.Fatal(err)
		}
		if lay.EngineID != "" {
			// 只对添加了引擎的分层作处理
			layers = append(layers, lay)
		}
	}
	fmt.Println("分层数据已获取")
	fmt.Println("开始获取存储源数据")
	var sources []metasource.MetaSourceDb
	bytes, err = esutil.SearchAll(common.DataBaseMetaData, common.TbMetaSource)
	if err != nil {
		fmt.Println("获取存储源数据失败:")
		log.Fatal(err)
	}
	data = esutil.GetSource(bytes)
	for _, v := range data {
		var sou metasource.MetaSourceDb
		err = json.Unmarshal([]byte(v), &sou)
		if err != nil {
			fmt.Println("解析存储源数据失败")
			log.Fatal(err)
		}
		sources = append(sources, sou)
	}
	fmt.Println("存储源数据已获取")
	fmt.Println("开始获取表数据")
	var tables []metasource.MetaSourceTb
	bytes, err = esutil.SearchAll(common.DataBaseMetaData, common.TbMetaTable)
	if err != nil {
		fmt.Println("获取表数据失败:")
		log.Fatal(err)
	}
	data = esutil.GetSource(bytes)
	for _, v := range data {
		var tb metasource.MetaSourceTb
		err = json.Unmarshal([]byte(v), &tb)
		if err != nil {
			fmt.Println("解析表数据失败")
			log.Fatal(err)
		}
		tables = append(tables, tb)
	}
	fmt.Println("表数据已获取")

	fmt.Println("开始修正分层数据...")
	for _, v := range layers {
		var sou metasource.MetaSourceDb
		for _, s := range sources {
			if s.LayerID == v.ID && !strings.Contains(s.SourceName, "wt") &&
				!strings.Contains(s.SourceName, "basedb") {
				sou = s
				break
			}
		}
		if sou.SourceName == "" {
			fmt.Printf("错误的分层：%#v\n", v)
			continue
		}
		if sou.SourceType == "hive" {
			v.DefaultSourceDb = sou.SourceName
		} else {
			v.DefaultSourceDb = "public"
		}
		v.DefaultDb = sou.SourceName
		v.DefaultSourceID = sou.ID
		doc := gin.H{
			"defaultdb":       v.DefaultDb,
			"defaultsourcedb": v.DefaultSourceDb,
			"defaultsourceid": v.DefaultSourceID,
			"dbexist":         true,
			"sourceexist":     true,
		}
		err = esutil.UpdByID(common.DataBaseMetaData, common.TbMetaLayer, v.ID, doc)
		if err != nil {
			fmt.Printf("分层更新失败:%s\n", v.Name)
			log.Fatal(err)
		}
	}
	fmt.Println("分层数据已修正")
	fmt.Println("开始修正存储源数据...")
	for _, v := range sources {
		v.DbName = v.SourceName
		if v.SourceType != "hive" {
			v.SourceName = "public"
		}
		doc := gin.H{
			"dbname":     v.DbName,
			"sourcename": v.SourceName,
		}
		err = esutil.UpdByID(common.DataBaseMetaData, common.TbMetaSource, v.ID, doc)
		if err != nil {
			fmt.Printf("分层更新失败:%s\n", v.DbName)
			log.Fatal(err)
		}
	}
	fmt.Println("存储源数据已修正")
	fmt.Println("开始修正表数据...")
	for _, v := range tables {
		v.DbName = v.SourceDb
		if v.DBType != "hive" {
			v.SourceDb = "public"
		}
		doc := gin.H{
			"dbname":   v.DbName,
			"sourcedb": v.SourceDb,
		}
		err = esutil.UpdByID(common.DataBaseMetaData, common.TbMetaTable, v.ID, doc)
		if err != nil {
			fmt.Printf("表更新失败:%s\n", v.DbName)
			log.Fatal(err)
		}
	}
	fmt.Println("表数据已修正")
	fmt.Println("所有数据已修正完成")
	return nil
}
