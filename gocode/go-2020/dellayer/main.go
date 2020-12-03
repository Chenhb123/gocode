package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("开始去除分层引擎信息...")
	now := time.Now()
	fmt.Println("开始时间：", now.Format("2006-01-02 15:04:05"))
	// 上级索引
	top := "danastudio-asset"
	// 待清空索引列表
	list := []string{
		"db_layer",
		"db_newasset",
	}
	var command, res string
	var err error
	for _, v := range list {
		fmt.Printf("开始清空索引:%s/%s...\n", top, v)
		command = fmt.Sprintf("curl -X DELETE '127.0.0.1:10100/%s/%s/_query' -d '{}'", top, v)
		res, err = ExeCommand(command)
		fmt.Println("执行结果：", res)
		if err != nil {
			fmt.Printf("删除索引%s/%s数据报错:%v\n", top, v, err.Error())
			continue
		}
		fmt.Printf("清空索引:%s/%s完成\n", top, v)
	}
	end := time.Now()
	fmt.Println("清空资产概览错误数据完成")
	fmt.Println("结束时间：", end.Format("2006-01-02 15:04:05"))
	consume := end.Sub(now)
	fmt.Println("总耗时：", consume)
}

// ExeCommand .
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
		return string(bytesErr[:]), err
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
