package main

import (
	"bytes"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

// Test .
type Test struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	ts := []Test{}
	var t Test
	for i := 5; i < 7; i++ {
		t.Name = fmt.Sprintf("%d-%d", i, i)
		t.Age = i + 1
		ts = append(ts, t)
	}
	for i := 0; i < len(ts); i++ {
		var es Test
		byteBody, _ := json.Marshal(ts[i])
		err := json.Unmarshal(byteBody, &es)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(es.Age)
		ts[i].ID = GenerateSubID()
	}
	var body string
	for _, v := range ts {
		byteBody, _ := json.Marshal(v)
		id := gjson.Get(string(byteBody), "id").String()
		fmt.Printf("%#v\n", id)
		fmt.Println(v)
		body = fmt.Sprintf(`%s
				{"create":{"_index":"%s","_type":"%s", "_id": "%s"}}
				%s`,
			body, "danastudio-asset", "tb_addrecord", id, string(byteBody))
	}
	fmt.Println(body)
	fmt.Println("1-----------1")
	if body != `` {
		body = fmt.Sprintf("%s\n", body)
		uri := fmt.Sprintf("/_bulk")
		_, err := Post(uri, body)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(respones))
	}
	fmt.Println("end--------------end")
}

// Idrand 生成大写字母+数字
func Idrand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 65}}, make([]byte, size)
	isall := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isall { // random ikind
			ikind = rand.Intn(2)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

// Post .
func Post(uri string, body string) (response []byte, err error) {

	url := fmt.Sprintf("http://%s:%d%s",
		"192.168.2.80",
		10100,
		uri)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return []byte(""), err
	}

	req.SetBasicAuth("admin", "admin123456")

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()

	//索引刷新
	err = FlushEs("192.168.2.80", 10100, "danastudio-asset")
	if err != nil {
		return []byte(""), err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil
}

// FlushEs .
func FlushEs(ip string, port int, index string) (err error) {
	client := &http.Client{}
	//索引刷新
	nurl := fmt.Sprintf("http://%s:%d/%s/_refresh",
		ip,
		port,
		index)
	nreq, err := http.NewRequest("POST", nurl, nil)
	if err != nil {
		return err
	}

	nreq.SetBasicAuth("admin", "admin123456")

	nresp, err := client.Do(nreq)
	if err != nil {
		return err
	}
	defer nresp.Body.Close()
	return nil
}

// GetGUID get guid.
func GetGUID() string {
	b := make([]byte, 20)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

// GetMd5String get md5string.
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-")
var n = 20

// GenerateSubID .
func GenerateSubID() string {
	b := make([]rune, n)
	for i := range b {
		index, _ := crand.Int(crand.Reader, big.NewInt(int64(len(letterRunes))))
		ind := int(index.Int64())
		b[i] = letterRunes[ind]
	}
	return string(b)
}
