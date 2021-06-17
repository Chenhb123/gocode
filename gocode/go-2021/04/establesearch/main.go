package main

import (
	"datatom.com/metadata/common"
	"datatom.com/metadata/httpdo/esutil"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	"os"
)

// MetaSourceTb 元数据开发--存储源下的表结构
type MetaSourceTb struct {
	ID         string `json:"id"`                // 元数据表id
	LayerName  string `json:"layername"`         // 层名称
	LayerID    string `json:"layerid"`           // 层id
	EngineID   string `json:"engineid"`          // 引擎id
	DbName     string `json:"dbname"`            // 数据库名称
	SourceDb   string `json:"sourcedb"`          // 存储源名称,对hive即为数据库名称,对stork/teryx指schema
	SourceID   string `json:"sourceid"`          // 存储源id
	Schema     string `json:"schema"`            // stork模式
	DBType     string `json:"dbtype"`            // 引擎类型 hive/stork
	FileFormat string `json:"fileformat"`        // 文件格式 Orcfile/Textfile
	TBName     string `json:"tbname"`            // 表名称
	Chinese    string `json:"chinese,omitempty"` // 表中文名
	Describe   string `json:"describe"`          // 表描述
	TBType     string `json:"tbtype"`            // 表类型 未分类/事实表/代码表
	IsRT       bool   `json:"isrt"`              // 是否实时 是则true
	Location   string `json:"location"`          // 表底层存储位置
	CatalogID  string `json:"catalogid"`         // 所属最底层编目id
	ProjectID  string `json:"projectid"`         // 所属项目id
	UserID     string `json:"userid"`            // 用户id
	Username   string `json:"username"`          // 用户名
	Created    string `json:"created"`           // 创建时间
	Modified   string `json:"modified"`          // 修改时间
	Updated    string `json:"updated"`

	StorageSpace string `json:"storagespace,omitempty"`
	SizeByte     uint64 `json:"sizebyte,omitempty"`
	UpdateSpace  uint64 `json:"updatespace,omitempty"`
	RowsNum      int    `json:"rowsnum,omitempty"`
	UpdateNum    int    `json:"updateNum,omitempty"`
	SyncTime     string `json:"synctime,omitempty"` //数据源同步时间
	SyncFreq     string `json:"syncfreq,omitempty"`
	VersionNum   int    `json:"versionnum"` // 版本号
}

func main() {
	fmt.Println("命令行参数长度：", len(os.Args))
	// 遍历 os.Args 切片，就可以得到所有的命令行输入参数值
	if len(os.Args) != 2 {
		log.Fatal("参数格式不正确，示例：./searchtask \"tableid\"")
	}
	tableid := os.Args[1]
	var sourcetb MetaSourceTb
	sourcestr, err := esutil.SearchByID("danastudio_metadata", "meta_sourcetb", tableid)
	if err != nil {
		log.Fatalf("获取表信息失败:%s", err.Error())
	}
	err = json.Unmarshal([]byte(sourcestr), &sourcetb)
	if err != nil {
		log.Fatalf("解析表数据失败:%s", err.Error())
	}
	editingTable(sourcetb)
}

func editingTable(info MetaSourceTb) (bool, error) {
	var target bool
	// 判定表是否被单例抽取任务依赖
	var musts []interface{}
	var term gin.H
	term = gin.H{
		"term": gin.H{
			"sourcedb.odsdatabase": info.DbName,
		},
	}
	musts = append(musts, term)
	term = gin.H{
		"term": gin.H{
			"sourcedb.table": info.TBName,
		},
	}
	musts = append(musts, term)
	if info.DBType == common.STORK || info.DBType == common.TERYX || info.DBType == common.GAUSSDB || info.DBType == common.DMDB {
		term = gin.H{
			"term": gin.H{
				"sourcedb.schema": info.SourceDb,
			},
		}
		musts = append(musts, term)
	}
	body := gin.H{
		"query": gin.H{
			"bool": gin.H{
				"must": musts,
			},
		},
	}
	bytes, err := esutil.SearchByTerm(common.DataBaseLab, common.TbAccess, body)
	if err != nil {
		//logger.Error.Printf("查询表抽取任务依赖失败:%s", err.Error())
		return target, err
	}
	accessStr := esutil.GetSource(bytes)
	for _, v := range accessStr {
		name := gjson.Get(v, "notename").String()
		fmt.Println("采集任务:", name)
	}
	// 判定表是否被批量抽取任务依赖
	musts = make([]interface{}, 0)
	term = gin.H{
		"term": gin.H{
			"sourcedb.odsdatabase": info.DbName,
		},
	}
	musts = append(musts, term)
	term = gin.H{
		"term": gin.H{
			"extractdb.batchtables": info.TBName,
		},
	}
	musts = append(musts, term)
	if info.DBType == common.STORK || info.DBType == common.TERYX || info.DBType == common.GAUSSDB {
		term = gin.H{
			"term": gin.H{
				"sourcedb.schema": info.SourceDb,
			},
		}
		musts = append(musts, term)
	}
	body = gin.H{
		"query": gin.H{
			"bool": gin.H{
				"must": musts,
			},
		},
	}
	bytes, err = esutil.SearchByTerm(common.DataBaseLab, common.TbAccess, body)
	if err != nil {
		//logger.Error.Printf("查询表抽取任务依赖失败:%s", err.Error())
		return target, err
	}
	accessStr = esutil.GetSource(bytes)
	for _, v := range accessStr {
		name := gjson.Get(v, "notename").String()
		fmt.Println("采集任务:", name)
	}
	// 判定是否被治理任务依赖
	GovernTaskDependency(info.DbName, info.DBType, info.SourceDb, info.TBName, info.EngineID)

	return true, nil
}

// GovernTaskDependency 判定表/存储源是否被治理任务依赖,true->被任务依赖
func GovernTaskDependency(dbname, dbtype, schema, tbname, engineid string) (bool, int, error) {
	if engineid == "" || dbname == "" || dbtype == "" {
		return false, 0, fmt.Errorf("判定表/存储源是否被治理任务依赖失败,参数异常:%s,%s,%s",
			engineid, dbname, dbtype)
	}
	var target bool
	var total int
	// 是否被表输入依赖
	musts := governQuery("tbinfo", dbname, dbtype, schema, tbname, engineid)
	body := gin.H{
		"query": gin.H{
			"bool": gin.H{
				"must": musts,
			},
		},
		"size": common.UnLimitSize,
	}
	bytes, err := esutil.SearchByTerm(common.DataBaseLab, common.TbDev, body)
	if err != nil {
		return target, 0, err
	}
	devstr := esutil.GetSource(bytes)
	for _, v := range devstr {
		dags := gjson.Get(v, "dags").Array()
		for _, d := range dags {
			name := d.Get("dagname").String()
			fmt.Println("治理任务：", name)
		}
	}
	// 是否被中间表依赖
	musts = governQuery("newtbinfo", dbname, dbtype, schema, tbname, engineid)
	body = gin.H{
		"query": gin.H{
			"bool": gin.H{
				"must": musts,
			},
		},
		"size": common.UnLimitSize,
	}
	bytes, err = esutil.SearchByTerm(common.DataBaseLab, common.TbDev, body)
	if err != nil {
		return target, 0, err
	}
	devstr = esutil.GetSource(bytes)
	for _, v := range devstr {
		dags := gjson.Get(v, "dags").Array()
		for _, d := range dags {
			name := d.Get("dagname").String()
			fmt.Println("治理任务：", name)
		}
	}
	// 是否被表输出依赖
	musts = governQuery("finaltable", dbname, dbtype, schema, tbname, engineid)
	body = gin.H{
		"query": gin.H{
			"bool": gin.H{
				"must": musts,
			},
		},
		"size": common.UnLimitSize,
	}
	bytes, err = esutil.SearchByTerm(common.DataBaseLab, common.TbDev, body)
	if err != nil {
		return target, 0, err
	}
	devstr = esutil.GetSource(bytes)
	for _, v := range devstr {
		dags := gjson.Get(v, "dags").Array()
		for _, d := range dags {
			name := d.Get("dagname").String()
			fmt.Println("治理任务：", name)
		}
	}
	return target, total, nil
}

func governQuery(query, dbname, dbtype, schema, tbname, engineid string) []interface{} {
	var musts []interface{}
	term := gin.H{
		"term": gin.H{
			query + ".dbname": dbname,
		},
	}
	musts = append(musts, term)
	term = gin.H{
		"term": gin.H{
			query + ".engineid": engineid,
		},
	}
	musts = append(musts, term)
	if dbtype == common.STORK || dbtype == common.TERYX ||
		dbtype == common.GAUSSDB || dbtype == common.DMDB {
		term := gin.H{
			"term": gin.H{
				query + ".schema": schema,
			},
		}
		musts = append(musts, term)
	}
	if tbname != "" {
		term := gin.H{
			"term": gin.H{
				query + ".tbname": tbname,
			},
		}
		musts = append(musts, term)
	}
	return musts
}
