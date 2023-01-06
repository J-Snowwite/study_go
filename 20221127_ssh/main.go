package main

import (
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	ipAddress := "1.1.1.1"
	userName := "root"
	passWord := "123456"

	client, _ := ssh.Dial("tcp", ipAddress, &ssh.ClientConfig{
		User:            userName,
		Auth:            []ssh.AuthMethod{ssh.Password(passWord)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	session, _ := client.NewSession()
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	session.RequestPty("linux", 32, 160, modes)
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	session.Shell()
	session.Wait()

}
