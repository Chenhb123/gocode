package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Unknwon/goconfig"

	"datatom.com/asset/common"
)

func main() {
	// WriteDefaultSourceCfg()
	cfg, err := goconfig.LoadConfigFile("./source_config.cfg")
	if err != nil {
		log.Fatal(err)
	}
	cycle := cfg.MustValue("config", "cycle")
	fmt.Println(cycle)
	starttime := cfg.MustInt("config", "starttime")
	fmt.Println(starttime)
	onoff := cfg.MustBool("config", "onoff")
	fmt.Println(onoff)
}

// PathExists 判断路径所在文件/文件夹是否存在, true-存在，false-不存在
// 如果返回的错误为nil,说明文件或文件夹存在,
// 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
// 如果返回的错误为其它类型,则不确定是否在存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		// 存在
		return true, nil
	}
	if os.IsNotExist(err) {
		// 不存在
		return false, nil
	}
	// 不确定
	return false, err
}

// WriteDefaultSourceCfg 写入默认的资产概览配置信息
func WriteDefaultSourceCfg() error {
	target, err := PathExists("./source_config.cfg")
	if err != nil {
		return err
	}
	if target {
		// 文件已存在，直接返回
		return nil
	}
	// 写入
	file, err := os.OpenFile("./source_config.cfg", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(common.SourceConfig)
	return err
}
