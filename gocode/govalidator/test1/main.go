package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"datatom.com/ants/httpdo"
	"golang.org/x/crypto/ssh"
)

func main() {
	exec := fmt.Sprintf("cd /root && python hdfs_ip.py")
	buf, err := httpdo.RemoteLogin("root", "datatom", "192.168.2.70", exec, 22)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ip := strings.TrimSpace(string(buf[:]))
	reg := regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+`)
	res := reg.FindAllString(ip, -1)
	if len(res) != 1 {
		fmt.Printf("请选择高可用的hdfs集群: %s\n", "192.168.2.103")
	}
	fmt.Println(res[0])
}

func RemoteLogin(user, password, host, exec string, port int) (string, error) {
	//fmt.Println("into remotelogin")
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
	}
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return "ssh err", err
	}
	defer client.Close()
	// create session
	if session, err = client.NewSession(); err != nil {
		return "create session err", err
	}
	defer session.Close()
	tempExec := "cd /root && ls"
	_, err = session.Output(tempExec)
	if err != nil {
		return "exec tempCmd error", err
	}
	buf, err := session.Output(exec)
	if err != nil {
		return "exec cmd err", err
	}
	return string(buf), nil
}
