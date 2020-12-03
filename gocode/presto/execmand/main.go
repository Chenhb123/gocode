package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	cmd := `python tdh_hive.py 192.168.90.21 10000 hive 123456 "select * from "dwddb"."dwd_person_sc" limit 100"`
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

// ExeCommand .
func ExeCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	fmt.Println("1---------------1")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	fmt.Println("122---------------1")
	if err := cmd.Start(); err != nil {
		return "", err
	}
	fmt.Println("1888888888---------------1")
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", err
	}
	fmt.Println("14333---------------1")
	if len(bytesErr) != 0 {
		return string(bytesErr[:]), err
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	fmt.Println("1666666666666---------------1")
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	fmt.Println("1777777777---------------1")
	i := (string(bytes[:]))
	return i, nil
}
