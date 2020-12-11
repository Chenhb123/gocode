package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var AuthUser = "eagles"
var AuthPasswd = "datatom.com"
var IP = "127.0.0.1"
var Port = "10100"
var Uri = "/_cat/indices"

func main() {
	nodes := []string{"192.168.70.77", "192.168.70.62", "192.168.70.63"}
	var err error
	node := nodes[0]
	nodes = nodes[1:]
	for {
		fmt.Println("node:", node)
		_, err = Post(node, Port, Uri, "")
		if err != nil {
			fmt.Println("err:", err.Error())
			circle([]string{}, nodes)
		}
		fmt.Println("IP:", IP)
		//break
		time.Sleep(30 * time.Second)
		node = IP
	}
}

func circle(usedNodes, nodes []string) {
	if len(usedNodes) == len(nodes) {
		// 轮询结束也未找到可用节点，休眠一定时间后重新查询
		// 恢复原始服务节点
		IP = "127.0.0.1"
		// 休眠
		time.Sleep(10 * time.Minute)
		// 重新开始
		usedNodes = make([]string, 0)
	}
	if len(nodes) == 0 {
		nodes = append(nodes, "127.0.0.1")
	}
	if len(usedNodes) == 0 {
		usedNodes = append(usedNodes, nodes[0])
	} else {
		usedNodes = append(usedNodes, nodes[len(usedNodes)])
	}
	node := usedNodes[len(usedNodes)-1]
	bytes, err := Post(node, Port, Uri, "")
	if err != nil {
		fmt.Printf("节点%s未连通:%s,详细报错信息:%s\n", node, err.Error(), string(bytes))
		circle(usedNodes, nodes)
		return
	}
	// 未报错则返回
	IP = usedNodes[len(usedNodes)-1]
	return
}

func Post(ip, port, uri string, body string) (response []byte, err error) {

	url := fmt.Sprintf("http://%s:%s%s",
		ip,
		port,
		uri)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return []byte(""), err
	}

	req.SetBasicAuth(AuthUser, AuthPasswd)

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil
}
