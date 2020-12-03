package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "os"
	"flag"
	"strconv"
	"strings"

	// "time"

	"datatom.com/ants/common"
	// "datatom.com/ants/httpclient"
	// "datatom.com/ants/httpdo"
	// "datatom.com/ants/logger"
	// "datatom.com/ants/source"
	tools "datatom.com/tools/httpdo"
	// "datatom/gin.v1"
	// simplejson "github.com/bitly/go-simplejson"
	hive "github.com/dazheng/gohive"
	"github.com/go-ini/ini"
	// "github.com/go-ini/ini"
	// "github.com/go-xorm/xorm"
	// "github.com/pkg/sftp"
	// "github.com/tidwall/gjson"
	// "github.com/tidwall/sjson"
)

const (
	DefMode    = "common"
	DefPresto  = "/etc/danastudio/presto_config.cfg"
	ModePresto = "presto"
)

func main() {
	var hiveVersion string
	var odsip string
	var odsport string
	var odsuser string
	var odspassword string
	var tableSqltxt string
	var token string
	var prestoPath string
	var prestouser string
	var cdhIP string
	var serviceuser string
	var loginPasswd string
	var kinitPath string
	var execMode string   //执行模式,common 为按引擎版本区分，presto 为直接使用 presto
	var prestoConf string //presto 配置文件路径
	var dbName string     //presto 的数据库
	var showRes bool      //是否展示执行结果，默认为 false

	// /etc/danastudio/hiveExec -hiveversion=%s -ip=%s -port=%s -odsuser=%s -odsuserpasswd=%s -kinitpath=%s
	// -prestopath=%s -prestouser=%s -loginpasswd=%s -serviceuser=%s -cdhip=%s -tbsql="%s" -mode="%s" -prestoconf="%s -dbname="%s" -showres=%s

	// /etc/danastudio/hiveExec -mode=presto -showres=true -prestoconf=/etc/danastudio/presto_config.cfg -tbsql="show tables"

	flag.StringVar(&hiveVersion, "hiveversion", "", "")
	flag.StringVar(&odsip, "ip", "", "")
	flag.StringVar(&odsport, "port", "", "")
	flag.StringVar(&odsuser, "odsuser", "", "")
	flag.StringVar(&odspassword, "odsuserpasswd", "", "")
	flag.StringVar(&prestoPath, "prestopath", "", "")
	flag.StringVar(&prestouser, "prestouser", "", "")
	flag.StringVar(&cdhIP, "cdhip", "", "")
	flag.StringVar(&serviceuser, "serviceuser", "", "")
	flag.StringVar(&loginPasswd, "loginpasswd", "", "")
	flag.StringVar(&kinitPath, "kinitpath", "", "")
	flag.StringVar(&tableSqltxt, "tbsql", "", "")
	flag.StringVar(&token, "token", "", "")
	flag.StringVar(&execMode, "mode", DefMode, "")
	flag.StringVar(&prestoConf, "prestoconf", DefPresto, "")
	flag.StringVar(&dbName, "dbname", "", "")
	flag.BoolVar(&showRes, "showres", false, "")

	flag.Parse()

	if execMode == ModePresto { //presto 执行模式
		if dbName == "" {
			fmt.Println("sql执行失败 : 数据库名为空")
			return
		}
		conn, err := PrestoConnect(dbName, prestoConf)
		if err != nil {
			fmt.Println("sql执行失败 : presto 连接失败 " + err.Error())
			return
		}
		defer conn.Close()
		rows, err := conn.Query(tableSqltxt)
		if err != nil {
			fmt.Println("sql执行失败 : presto 执行失败 " + err.Error())
			return
		}
		if showRes {
			resMap, _ := sqlrows2Maps(rows)
			resByte, _ := json.Marshal(resMap)
			fmt.Println(string(resByte))
		}
		return
	}
	// 获取节点端口号、用户
	sshports, sshusers := tools.HiveSSHInfo()
	hiveport, hiveuser := sshports["hive"], sshusers["hive"]
	switch hiveVersion {
	case "", "3.0", common.Hive12, common.Hive3:
		str := odsip + ":" + odsport
		var db *hive.Connection
		db, err := hive.ConnectWithUser(str, odsuser, odspassword, hive.DefaultOptions)
		if err != nil {
			return
		}
		defer db.Close()
		_, err = db.Exec(tableSqltxt)
		if err != nil {
			fmt.Println("sql执行失败: ", str, tableSqltxt)
			return
		}
	case common.HiveTdh:
		var port int
		port, err := strconv.Atoi(odsport)
		if err != nil {
			return
		}
		_, err = tools.TDHiveRes(tableSqltxt, odsip, odsuser, odspassword, port)
		if err != nil {
			fmt.Println("sql执行失败: ", odsip, odsport, tableSqltxt)
			return
		}
	case common.HiveTDC:
		var port int
		port, err := strconv.Atoi(odsport)
		if err != nil {
			return
		}
		_, err = tools.TDCRes(odsip, port, tableSqltxt, token)
		if err != nil {
			fmt.Println("sql执行失败: ", odsip, odsport, tableSqltxt)
			return
		}
	case common.HiveCDH:
		exec := fmt.Sprintf("kinit -kt %s %s && beeline -u \"jdbc:hive2://%s:%v/default;principal=%s\" --silent=true --outputformat=csv2 -e \"%s\"",
			prestoPath, prestouser, cdhIP, odsport, serviceuser, tableSqltxt)
		var res string
		res, err := tools.ExecCmd(odsip, hiveport, hiveuser, loginPasswd, exec)
		if err != nil {
			return
		}
		if strings.Contains(res, "Error") {
			fmt.Println("sql执行失败: ", odsip, odsport, tableSqltxt)
			return
		}
	case common.HiveHuawei:
		exec := fmt.Sprintf("source %s/bigdata_env && kinit -kt %s %s && beeline --silent=true --outputformat=csv2 -e \"%s\"",
			kinitPath, prestoPath, prestouser, tableSqltxt)
		var res string
		res, err := tools.ExecCmd(odsip, hiveport, hiveuser, loginPasswd, exec)
		if err != nil {
			return
		}
		if strings.Contains(res, "Error") {
			return
		}
	default:
		fmt.Println("未知的hive版本：version: %s, odsip: %s", hiveVersion, odsip)
		return
	}
	fmt.Println("sql执行成功", tableSqltxt)
}

func PrestoConnect(database string, confPath string) (*sql.DB, error) {
	//s := "http://" + sp.User + "@" + sp.Ip + ":" + strconv.Itoa(sp.Port) + "?catalog=hive&schema=" + database

	//s = "http://root@192.168.2.182:9988?catalog=hive&schema=test"
	//通过配置的presto获取连接信息
	//---------------------------------------------------------------------------------------
	cfg, err := ini.Load(confPath)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}
	secServer, _ := cfg.GetSection("config")
	keyip, _ := secServer.GetKey("presto_ip")
	ip := keyip.Value()
	keyPort, _ := secServer.GetKey("presto_port")
	port := keyPort.Value()
	s := "http://root@" + ip + ":" + port + "?catalog=hive&schema=" + database
	//fmt.Println(s)
	//fmt.Println("打印默认连接信息")
	//---------------------------------------------------------------------------------------
	db, err := sql.Open("presto", s)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("连接失败")
		return nil, err
	}
	return db, nil
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
