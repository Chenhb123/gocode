package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

// Res1 。
var Res1 map[string]map[string]map[string]string

func main() {
	sqltxts := []string{}
	sqltxts = append(sqltxts, "show databases")
	res, err := TDHiveRes(sqltxts, "192.168.2.172", "10000", "hive", "123456")
	if err != nil {
		log.Fatal(err)
	}
	js, err := simplejson.NewJson([]byte(res))
	if err != nil {
		log.Fatal(err)
	}
	rec1, err := js.Get("0").String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rec1)
}

// TDHiveRes 连接hive并返回sql执行结果
func TDHiveRes(sqltxts []string, ip, port, user, passwd string) (string, error) {
	var res string
	if len(sqltxts) == 0 || ip == "" || port == "" || user == "" || passwd == "" {
		return res, fmt.Errorf("参数错误：ip: %s, port: %s, user: %s, passwd: %s\n sqltxts: %v",
			ip, port, user, passwd, sqltxts)
	}
	fileName := "/etc/danastudio/tdhsqls.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return res, err
	}
	defer file.Close()
	sqls := ""
	for _, v := range sqltxts {
		sqls += fmt.Sprintf("%s;\n", v)
	}
	if _, err = file.WriteString(sqls); err != nil {
		return res, err
	}
	command := "cat /etc/danastudio/tdhsqls.txt"
	res, err = ExeCommand(command)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	fmt.Println("file content:\n", res)
	command = fmt.Sprintf("cd %s && python %s %s %s %s %s %s",
		"/etc/danastudio/", "go_tdh.py", ip, port, user, passwd, "tdhsqls.txt")
	fmt.Println("command:", command)
	res, err = ExeCommand(command)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	fmt.Println("res:", res)
	if strings.Contains(strings.ToLower(res), "error") {
		return res, fmt.Errorf("执行出错：sqltxt: %s\n, res: %s", sqls, res)
	}
	return res, nil
}

// ExeCommand linux操作
func ExeCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", err
	}

	if len(bytesErr) != 0 {
		return "", err
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	i := (string(bytes[:]))
	return i, nil
}
