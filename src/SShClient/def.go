package SshClient

import (
	"io"
	"log"

	"golang.org/x/crypto/ssh"
)

type ClientInfo struct {
	Ipaddr     string
	Console    string
	Name       string
	Login      string
	Passwd     string
	Routertype int // CISCO/FLEXSWITCH etc

	client     *ssh.Client // This will contains client to remote switch
	r          io.Reader
	w          io.Writer
	subprocess *ssh.Session // Create a new subprocess on switch to receive config
	WriteCh    chan string  // Write to this will go subprocess
	ReadCh     chan string  // we will wait till we get response
	log        *log.Logger
}

type SshClientInfo struct {
	// Name of Resource and all the informatin
	// of client
	SshInfo map[string]*ClientInfo
}
