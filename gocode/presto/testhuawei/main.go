package main

import (
	// "context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	// "github.com/beltran/gohive"
)

func main() {
	command := "ssh 192.168.2.80 'cd ~ && df -h'"
	_, err := ExeCommand(command)
	if err != nil {
		fmt.Println("1------------1")
		log.Fatal(err)
	}
	// ctx := context.Background()
	// configuration := gohive.NewConnectConfiguration()
	// // configuration.TransportMode = "binary"
	// // configuration.Service = "hive/hadoop.hadoop.com"
	// // configuration.Service = "hive"
	// configuration.Username = "hive"
	// configuration.Password = "123456"
	// conn, err := gohive.Connect("192.168.90.105", 10000, "NOSASL", configuration)
	// if err != nil {
	// 	fmt.Println("2--------------2")
	// 	log.Fatal(err)
	// }
	// defer conn.Close()
	// cursor := conn.Cursor()
	// defer cursor.Close()
	// cursor.Exec(ctx, "show databases")
	// if cursor.Err != nil {
	// 	log.Fatal(cursor.Err)
	// }
	// for cursor.HasMore(ctx) {
	// 	if cursor.Err != nil {
	// 		log.Fatal(cursor.Err)
	// 	}
	// 	mapRes := cursor.RowMap(ctx)
	// 	if cursor.Err != nil {
	// 		log.Fatal(cursor.Err)
	// 	}
	// 	fmt.Println("mapRes:", mapRes)
	// }
}

// ExeCommand 。
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
