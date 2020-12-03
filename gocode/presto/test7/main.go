package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"datatom.com/ants/logger"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	_, err := localCopys("root", "datatom", "192.168.2.70", "/home/", "/home/hdfs_ip.py")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("end")
}

func localCopys(user, password, host, localpath, remotepath string) (string, error) {
	var (
		err        error
		sftpClient *sftp.Client
	)
	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err = connect(user, password, host, 22)
	if err != nil {
		return "传输失败", err
	}
	defer sftpClient.Close()
	srcFile, err := os.Open(localpath)
	if err != nil {
		logger.Error.Println(err)
		return "传输失败", err
	}
	defer srcFile.Close()
	var remoteFileName = path.Base(localpath)
	// /home/hdfs_ip.py
	dstFile, err := sftpClient.Create(path.Join(remotepath, remoteFileName))
	if err != nil {
		logger.Error.Println(err)
		return "传输失败", err
	}
	defer dstFile.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf[0:n])
	}
	return "传输成功", nil
}

//远程copy文件
func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
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

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	defer sshClient.Close()
	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	fmt.Println("11111111111111111")
	return sftpClient, nil
}
