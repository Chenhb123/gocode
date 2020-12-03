package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

func main() {
	readBufio()
}

func readBufio() error {
	file, err := os.Open(`E:\share\measurements.csv`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// os.O_CREATE os.O_APPEND os.O_WRONLY
	fmtFile, err := os.OpenFile(`E:\share\mes_fmt.csv`, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	errFile, err := os.OpenFile(`E:\share\mes_err.csv`, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)
	var n int
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		res := string(line)
		res, err = convertEncoding(res, "GBK", "UTF-8")
		if err != nil {
			log.Fatal(err)
		}
		resArr := strings.Split(res, ",")
		if len(resArr) >= 2 && len(resArr[1]) == 8 {
			_, err = fmtFile.WriteString(res + "\n")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err = errFile.WriteString(res + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		n++
	}
	fmt.Println("n:", n)
	return nil
}

// data: 待转码的数据
// src: 源数据编码
// dst: 目标数据编码
func convertEncoding(data, src, dst string) (string, error) {
	var res string
	srcDec := mahonia.NewDecoder(src)
	dstDec := mahonia.NewDecoder(dst)
	srcData := srcDec.ConvertString(data)
	_, resByte, err := dstDec.Translate([]byte(srcData), true)
	if err != nil {
		return res, err
	}
	res = string(resByte)
	return res, nil
}
