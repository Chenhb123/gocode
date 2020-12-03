package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bitly/go-simplejson"
)

func main() {
	fmt.Println("命令行参数长度：", len(os.Args))
	// 遍历 os.Args 切片，就可以得到所有的命令行输入参数值
	if len(os.Args) != 6 {
		log.Fatal("参数格式不正确，示例：./createuser.exe \"excel.xlsx\" \"B\", \"192.168.2.12\" admin admin")
	}
	ip, user, passwd := os.Args[3], os.Args[4], os.Args[5]
	body := fmt.Sprintf(`{"username":"%s","password":"%s"}`, user, passwd)
	res, err := PostAuthorization(ip, "/api/pandabi/user/login", body, "json", "")
	if err != nil {
		log.Fatal(err)
	}
	js, err := simplejson.NewJson(res)
	if err != nil {
		log.Fatal(err)
	}
	token, err := js.Get("token").String()
	if err != nil {
		log.Fatal(err)
	}
	token = fmt.Sprintf("Bearer %s", token)
	fmt.Println(token)
	// Get value from cell by given worksheet name and axis.
	path := os.Args[1]
	column := os.Args[2]
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	var col string
	for i := 2; i < 10000; i++ {
		col = fmt.Sprintf("%s%d", column, i)
		cell, err := f.GetCellValue("Sheet1", col)
		if err != nil {
			fmt.Println(err)
			return
		}
		if cell == "" {
			continue
		}
		fmt.Println(i, ":", cell)
		body = fmt.Sprintf(`username=%s&password=123456`, cell)
		res, err = PostAuthorization(ip, "/api/pandabi/user/create", body, "urlencoded", token)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(res))
	}
	time.Sleep(10 * time.Second)
}

// PostAuthorization 需要带有token认证的请求
func PostAuthorization(ip, uri, body, style, authorization string) (response []byte, err error) {

	url := fmt.Sprintf("http://%s%s",
		ip,
		uri)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return []byte(""), err
	}

	req.Header.Set("Content-Type", "application/json")
	if style == "urlencoded" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if authorization != "" {
		req.Header.Set("Authorization", authorization)
	}
	//req.SetBasicAuth(AuthUser, AuthPasswd)

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil
}
