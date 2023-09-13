package task_ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os/exec"
	"time"
)

type SshServer struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Host          string `json:"host"`
	Port          int32
	AuthMode      string
	PrivateKey    string
	PrivateSecret string
	Cmd           string
	ClientConfig  *ssh.ClientConfig
	Session       *ssh.Session
	Err           error
}

func (s SshServer) ExecRemoteSsh() (string, error) {
	//privateKey := "\/++\nuoaG/"

	var authMethods []ssh.AuthMethod
	if s.Password != "" {
		authMethods = append(authMethods, ssh.Password(s.Password))
	}
	if s.PrivateKey != "" {
		hostSigner, err := ssh.ParsePrivateKey([]byte(s.PrivateKey))
		if err != nil {
			//log.Fatal(err)
			//err = fmt.Errorf("Failed to dial: " + err.Error())
			return "", err
		}
		//_ = hostSigner
		authMethods = append(authMethods, ssh.PublicKeys(hostSigner))
	}
	config := &ssh.ClientConfig{
		User:            s.Username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	//sshConfig.AddHostKey(hostSigner)
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port), config)
	if err != nil {
		//panic("Failed to dial: " + err.Error())
		err = fmt.Errorf("Failed to dial: " + err.Error())
		return "", err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		err = fmt.Errorf("Failed to new session: " + err.Error())
		//panic("Failed to new session: " + err.Error())
		return "", err
	}
	defer session.Close()
	var exitCode = 0
	output, err := session.Output(s.Cmd)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		}
		err = fmt.Errorf("Failed to session output: " + err.Error())
		return "", err
	}

	//fmt.Println("output:", string(output))
	now := time.Now().Local().Format("2006-01-02 15:04:05")
	resp := string(output) + fmt.Sprintf("\n%s\n%s Task Run Finished, Return exitCode:%5d", "", now, exitCode)
	return resp, nil
}
