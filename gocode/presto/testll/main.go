package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

// ls 输出结果

func main() {
	existcmd := fmt.Sprintf(`cd /opt/danastudio && ls`)
	existfile, err := ExeCommand(existcmd)
	if err != nil {
		fmt.Println("err.........")
		log.Fatal(err)
	}
	// fmt.Println(existfile)
	ss := strings.Split(existfile, "\n")
	// fmt.Println(ss)
	for _, v := range ss {
		if strings.Contains(v, "ants") {
			fmt.Println(v)
		}
	}
}

// ExeCommand .
func ExeCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	fmt.Println("6-------------6")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	fmt.Println("5-------------5")
	if err := cmd.Start(); err != nil {
		return "", err
	}
	fmt.Println("4--------------4")
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", err
	}
	fmt.Println("3--------------3")
	fmt.Println("err:", string(bytesErr))
	if len(bytesErr) != 0 && !strings.Contains(string(bytesErr), "Xferd") {
		return string(bytesErr[:]), err
	}
	fmt.Println("2-----------------2")
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	fmt.Println("1----------1")
	i := (string(bytes[:]))
	return i, nil
}
