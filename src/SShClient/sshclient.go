package SshClient

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

// Create Ssh client object that can be used
// in Resource manager etc to call this.
func Init() *SshClientInfo {
	fmt.Println("Creating Ssh Main init")
	return &SshClientInfo{SshInfo: make(map[string]*ClientInfo)}
}
func (sshclient *SshClientInfo) createClientInfo(mymap map[string]string) *ClientInfo {
	newclient := &ClientInfo{Login: "admin", Passwd: "mysnaproute"}
	if name, ok := mymap["NAME"]; ok {
		newclient.Name = name
	} else {
		fmt.Println("NAME NOT FOND IN CLEINT")
		return nil
	}

	if val, ok := mymap["IPADDR"]; ok {
		newclient.Ipaddr = val
	} else {
		fmt.Println("IP NOT FOND IN CLEINT")
		return nil
	}
	config := &ssh.ClientConfig{User: newclient.Login,
		Auth:            []ssh.AuthMethod{ssh.Password(newclient.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), Timeout: 0}
	// Create a new sshclient process to the remote switch
	client, err := ssh.Dial("tcp", newclient.Ipaddr+":22", config)
	if err != nil {
		fmt.Println("Connection to remote side failed", client)
	}
	newclient.client = client
	newclient.CreateSubProcess()
	return newclient
}

func (Client *ClientInfo) CreateSubProcess() {
	session, _ := Client.client.NewSession()
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	sshIn, _ := session.StdinPipe()
	Client.w = sshIn
	sshOut, _ := session.StdoutPipe()
	Client.r = sshOut

	if err := session.RequestPty("xterm", 0, 200, modes); err != nil {
		fmt.Println("Failed1")
		session.Close()
		return
	}

	if err := session.Shell(); err != nil {
		fmt.Println("Failed2")
		session.Close()
		return
	}
	Client.subprocess = session

	// Starting Go routine to take command
	go Client.SessionGoRoutine()

}

func (sshclient *SshClientInfo) CreateClients(obj interface{}) {

	fmt.Println("Creating client  all")
	switch obj.(type) {
	case map[string]string:
		mymap := obj.(map[string]string)

		// CAPITAL "NAME" is mandatory for any Cliewnt
		if name, ok := mymap["NAME"]; !ok {
			fmt.Println("ERROR: Name elemetn not found in map")
			return
		} else {
			// Check whether that Client is already created
			if _, ok := sshclient.SshInfo[name]; !ok {
				newclient := sshclient.createClientInfo(mymap)
				sshclient.SshInfo[name] = newclient

				fmt.Println("Creating client  ", sshclient)
			}
		}
	default:
		fmt.Println("Underfined type received")

	}
}

func (sshclient *SshClientInfo) GetClient(name string) {
}
