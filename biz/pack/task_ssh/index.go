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
	Client        *ssh.Client
	Session       *ssh.Session
	Err           error
}

func (s *SshServer) ExecRemoteSshInit() error {
	//privateKey := "\/++\nuoaG/"
	var err error
	var authMethods []ssh.AuthMethod
	if s.Password != "" {
		authMethods = append(authMethods, ssh.Password(s.Password))
	}
	if s.PrivateKey != "" {
		var hostSigner ssh.Signer
		if s.PrivateSecret != "" {
			hostSigner, err = ssh.ParsePrivateKeyWithPassphrase([]byte(s.PrivateKey), []byte(s.PrivateSecret))
		} else {
			hostSigner, err = ssh.ParsePrivateKey([]byte(s.PrivateKey))
		}
		if err != nil {
			return err
		}
		//_ = hostSigner
		authMethods = append(authMethods, ssh.PublicKeys(hostSigner))
	}
	config := &ssh.ClientConfig{
		User:            s.Username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	s.Client, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port), config)
	if err != nil {
		err = fmt.Errorf("Failed to dial: " + err.Error())
		return err
	}
	//defer s.Client.Close()
	//s.Session, err = s.Client.NewSession()
	//if err != nil {
	//	err = fmt.Errorf("Failed to new session: " + err.Error())
	//	return err
	//}
	////defer s.Session.Close()
	return nil
}

func (s *SshServer) ExecRemoteSshCmd(Cmd string) (string, error) {
	var exitCode = 0
	session, err := s.Client.NewSession()
	if err != nil {
		err = fmt.Errorf("Failed to new session: " + err.Error())
		return "", err
	}
	defer session.Close()

	output, err := session.Output(Cmd)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		}
		err = fmt.Errorf("Failed to session output: " + err.Error())
		return "", err
	}
	now := time.Now().Local().Format("2006-01-02 15:04:05")
	resp := string(output) + fmt.Sprintf("\n%s\n%s Task Run Finished, Return exitCode:%5d", "", now, exitCode)
	return resp, nil
}

func (s *SshServer) Close() {
	_ = s.Client.Close()
}
