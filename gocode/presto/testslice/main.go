package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	tools "datatom.com/tools/httpdo"
)

func main() {
	command := `beeline -u "jdbc:hive2://192.168.90.105:10000" --silent=true --outputformat=csv2 -e "show databases"`
	res, err := tools.ExecCmd("192.168.90.105", 22, "root", "datatom", command)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("res:", res)
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
