package main

import (
	// "context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	// "github.com/beltran/gohive"
)

func main() {
	command := "python /home/chb/tdc_hive.py 192.168.90.91 31213 cKjIZGNM8jGgqQrtss4f-WHVTFZ0.TDH \"show databases\""
	fmt.Println("command:", command)
	res, err := ExeCommand(command)
	if err != nil {
		fmt.Println("1------------1")
		log.Fatal(err)
	}
	fmt.Println(res)
}

// ExeCommand 。
func ExeCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	// fmt.Println("6----------------6")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	// fmt.Println("5---------------5")
	if err := cmd.Start(); err != nil {
		return "", err
	}
	// fmt.Println("4---------------4")
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", err
	}
	// fmt.Println("3------------3")
	if len(bytesErr) != 0 && !strings.Contains(string(bytesErr), "WARN") {
		return string(bytesErr[:]), err
	}
	// fmt.Println("2------------2")
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	// fmt.Println("1--------1")
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	// i := (string(bytes[:]))
	// fmt.Println(string(bytes))
	return string(bytes), nil
}
