package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"
	"time"
	// "golang.org/x/sys/unix"
	// "fmt"
)

type configBean struct {
	Test string `json:"test"`
}

type config struct {
	LastModify time.Time
	Data       configBean
}

// Config .
var Config *config

func main() {
	const testConst = "Hello Const!"
	// var value string
}

func loadConfig(path string) error {
	locker := new(sync.RWMutex)
	// 读取配置文件内容
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	// 读取文件属性
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	// 验证文件的修改时间
	if Config != nil && fileInfo.ModTime().Before(Config.LastModify) {
		return errors.New("no need update")
	}
	// 解析文件内容
	var configBean configBean
	err = json.Unmarshal(data, &configBean)
	if err != nil {
		return err
	}
	var config = config{
		LastModify: fileInfo.ModTime(),
		Data:       configBean,
	}
	// 重新赋值更新配置文件
	locker.Lock()
	Config = &config
	locker.Unlock()

	return nil
}
