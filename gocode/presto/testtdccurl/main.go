package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("start........")
	// existcmd := fmt.Sprintf(`curl -X GET "http://192.168.90.91:31573/webhdfs/v1/tmp/detuo?op=LISTSTATUS&guardian_access_token=cKjIZGNM8jGgqQrtss4f-WHVTFZ0.TDH"`)
	// sql := `
	// 	create table if not exists "odsdb.test0302" (
	// 	"col" string, "col1" string, "col2" string, "col3" string)
	// 	ROW FORMAT DELIMITED  FIELDS TERMINATED BY '\u0001'
	// `
	// 	sql := fmt.Sprintf(`load DATA inpath '/tmp/detuo/%s'  INTO TABLE "%s.%s"`,
	// "test0302", "odsdb", "test0302")
	// sql := `
	// create table if not exists "odsdb"."realtest" (
	// "col" string, "col1" string, "col2" string, "col3" string)
	// CLUSTERED BY (col) INTO 3 BUCKETS
	// STORED AS ORC
	// TBLPROPERTIES ('transactional'='true')`
	// sql := `INSERT INTO TABLE "odsdb.realtest" SELECT * FROM "odsdb.test0302"`
	// sql := `drop table "odsdb.test0302"`
	sql := `create database if not exists chb0311`
	existcmd := fmt.Sprintf(`python /etc/danastudio/tdc_hive.py 192.168.90.91 31213 %s "%s"`,
		"cKjIZGNM8jGgqQrtss4f-WHVTFZ0.TDH", sql)
	existfile, err := ExeCommand(existcmd)
	if strings.Contains(existfile, "NoneType") {
		fmt.Println("exists:", existfile)
	}
	if err != nil {
		fmt.Println("err.........")
		format := fmt.Sprintf("err:%#v", err)
		fmt.Println("format:", format)
		log.Fatal(err)
	}

}

// ExeCommand .
func ExeCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	// fmt.Println("6-------------6")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	// fmt.Println("5-------------5")
	if err := cmd.Start(); err != nil {
		return "", err
	}
	// fmt.Println("4--------------4")
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", err
	}
	// fmt.Println("3--------------3")
	// fmt.Println("err:", string(bytesErr))
	if len(bytesErr) != 0 &&
		!strings.Contains(string(bytesErr), "WARN") &&
		!strings.Contains(string(bytesErr), "Xferd") {
		return string(bytesErr[:]), err
	}
	// fmt.Println("2-----------------2")
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	// fmt.Println("000000000000")
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	// fmt.Println("1----------1")
	i := (string(bytes[:]))
	return i, nil
}
