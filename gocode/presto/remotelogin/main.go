package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	// exec := "cd /home/hadoopclient && source bigdata_env && beeline -e 'show databases'"
	exec := `beeline -u "jdbc:hive2://192.168.90.105:10000/default;principal=hive/dn9025@CDH.COM" --silent=true --outputformat=csv2 -e "desc formatted "chbtest"."shijian""`
	res, err := RemoteLogin("root", "datatom", "192.168.90.104", exec, 22)
	if err != nil {
		log.Fatal(err)
	}
	var result string
	slice := strings.Split(res, "\n")
	for _, v := range slice {
		if strings.Contains(v, ",numRows") {
			result = v
			break
		}
	}
	pre := strings.LastIndex(result, ",")
	result = result[pre+1:]
	result = strings.TrimSpace(result)
	fmt.Println(result)
}

// RemoteLogin .
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
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
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
	buf, _ := session.CombinedOutput(exec)
	return string(buf), nil
}
