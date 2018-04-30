package main

import (
	"Commondefs"
	"EgressPipe"
	"Mgmt"
	"Rsrc"
	"SshClient"
	"TestInfra"
	"fmt"
	"time"
)

func main() {

	// This will be used for sending reference of other package
	// to required packages
	var objPool []interface{}
	rsc := Rsrc.Init()

	// Create ssh client object
	//sshclient := SshClient.Init()
	sshclient := SshClient.Init()

	objPool = append(objPool, rsc)
	objPool = append(objPool, sshclient)

	rsc.SetRefs(objPool)

	// Run the Resource Go Routine
	go rsc.ResourceRoutine()
	// Create Ingress Channel
	GlobalIngressCh := make(chan interface{})
	// Create Mgmt Pipeline - User managerment, Resource Management etc
	MgmtLeftCh, _ := Mgmt.CreateManagementPipeLine(GlobalIngressCh)
	// THis Contains Basic Sanity, L2, L3, etcc
	MgmtLeftCh, _ = TestInfra.CreateIngressTestPipeLine(MgmtLeftCh)

	// THis pipeline decided on how many times, we have recycle the
	// tests on failure etc
	MgmtLeftCh, _ = EgressPipe.CreateEgressPipeLine(MgmtLeftCh)
	// Data from Outside world is fed here to this Ingress Channel. Data can
	// from Web server as HTTP Login Page Response
	//req := &Commondefs.Request{User: &Commondefs.UserInfo{"Raju", "PRaju"}}
	// Create Resource Object
	time.Sleep(2 * time.Second)

	rsc.FileReadCh <- struct{}{}
	resp := &Commondefs.LoginResponse{}
	go SendData(resp, GlobalIngressCh, MgmtLeftCh)

	job := &Commondefs.UserJob{UserName: "Raju"}
	go SendData(job, GlobalIngressCh, MgmtLeftCh)

	time.Sleep(40 * time.Second)
	// wait for ever
	//select {}
}
func SendData(data interface{}, rightCh chan interface{}, leftCh chan interface{}) {
	fmt.Println("SEnding Data onto Channel")
	rightCh <- data
	recvData := <-leftCh
	DecodeReceivedData(recvData)
}
func DecodeReceivedData(data interface{}) {
	switch data.(type) {
	case *Commondefs.Request:
		fmt.Println("Request Recevied")
	case *Commondefs.LoginResponse:
		fmt.Println("Login Reponse Recevied")
	case *Commondefs.HttpUserResponse:
		fmt.Println("Login Response Recevied")
	default:
		fmt.Println("UNDEFINED RESPONSE RECEIVED:")
	}
	return
}
