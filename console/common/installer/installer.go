/*
   @Author: 1usir
   @Description:
   @File: installer
   @Version: 1.0.0
   @Date: 2024/7/2 15:02
*/

package installer

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"strconv"
	"time"
)

type Installer struct {
	client *SSHClient
}

func NewInstaller() *Installer {
	return &Installer{}
}

// Login 登录SSH服务
func (this *Installer) Login(credentials *Credentials) error {
	var hostKeyCallback ssh.HostKeyCallback = nil

	// 检查参数
	if len(credentials.Host) == 0 {
		return errors.New("'host' should not be empty")
	}
	if credentials.Port <= 0 {
		return errors.New("'port' should be greater than 0")
	}
	if len(credentials.Password) == 0 && len(credentials.PrivateKey) == 0 {
		return errors.New("require user 'password' or 'privateKey'")
	}

	// 不使用known_hosts
	if hostKeyCallback == nil {
		hostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		}
	}

	// 认证
	var methods = []ssh.AuthMethod{}
	if credentials.Method == "user" {
		{
			var authMethod = ssh.Password(credentials.Password)
			methods = append(methods, authMethod)
		}

		{
			authMethod := ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
				if len(questions) == 0 {
					return []string{}, nil
				}
				return []string{credentials.Password}, nil
			})
			methods = append(methods, authMethod)
		}
	} else if credentials.Method == "privateKey" {
		var signer ssh.Signer
		var err error
		if len(credentials.Passphrase) > 0 {
			signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(credentials.PrivateKey), []byte(credentials.Passphrase))
		} else {
			signer, err = ssh.ParsePrivateKey([]byte(credentials.PrivateKey))
		}
		if err != nil {
			return fmt.Errorf("parse private key: %w", err)
		}
		authMethod := ssh.PublicKeys(signer)
		methods = append(methods, authMethod)
	} else {
		return errors.New("invalid method '" + credentials.Method + "'")
	}

	// SSH客户端
	if len(credentials.Username) == 0 {
		credentials.Username = "root"
	}
	var config = &ssh.ClientConfig{
		User:            credentials.Username,
		Auth:            methods,
		HostKeyCallback: hostKeyCallback,
		Timeout:         5 * time.Second, // TODO 后期可以设置这个超时时间
	}

	sshClient, err := ssh.Dial("tcp", credentials.Host+":"+strconv.Itoa(credentials.Port), config)
	if err != nil {
		return err
	}
	client, err := NewSSHClient(sshClient)
	if err != nil {
		return err
	}

	if credentials.Sudo {
		client.Sudo(credentials.Password)
	}

	this.client = client

	return nil
}

func (this *Installer) Install(params map[string]string) error {
	// 检查目标目录是否存在
	_, err := this.client.Stat("/opt/hfagent")
	if err != nil {
		err = this.client.MkdirAll("/opt/hfagent")
		if err != nil {
			return fmt.Errorf("create directory  '%s' failed: %w", "/opt/hfagent", err)
		}
	}
	//执行命令下载可执行程序
	_, _, _ = this.client.Exec("curl -fLJ https://github.com/4hfire/4hfire/releases/latest/download/hfa -o /opt/hfagent/hfa && chmod +x hfa")

	// 设置环境变量
	{
		_, _, _ = this.client.Exec("export HFC_ENDPOINT=" + params["endpoints"])
		_, _, _ = this.client.Exec("export HFC_NODEID=" + params["nodeId"])
		_, _, _ = this.client.Exec("export HFC_SECRET=" + params["secret"])
	}

	// 启动
	_, stderr, err := this.client.Exec("/opt/hfagent/hfa start")
	if err != nil {
		return fmt.Errorf("start edge node failed: %w", err)
	}

	if len(stderr) > 0 {
		return errors.New("start edge node failed: " + stderr)
	}

	return nil
}

// Close 关闭SSH服务
func (this *Installer) Close() error {
	if this.client != nil {
		return this.client.Close()
	}
	return nil
}

func (this *Installer) uname() (uname string) {
	var unameRetries = 3

	for i := 0; i < unameRetries; i++ {
		for _, unameExe := range []string{"uname", "/bin/uname", "/usr/bin/uname"} {
			uname, _, _ = this.client.Exec(unameExe + " -a")
			if len(uname) > 0 {
				return
			}
		}
	}

	return "x86_64 GNU/Linux"
}
