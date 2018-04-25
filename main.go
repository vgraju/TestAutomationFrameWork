package main

import (
	"Commondefs"
	"EgressPipe"
	"Mgmt"
	"Tests"
	"fmt"
	"time"
)

func main() {

	// Create Ingress Channel
	GlobalIngressCh := make(chan interface{})
	// Create Mgmt Pipeline - User managerment, Resource Management etc
	MgmtLeftCh, _ := Mgmt.CreateManagementPipeLine(GlobalIngressCh)
	// THis Contains Basic Sanity, L2, L3, etcc
	MgmtLeftCh, _ = Tests.CreateIngressTestPipeLine(MgmtLeftCh)

	// THis pipeline decided on how many times, we have recycle the
	// tests on failure etc
	MgmtLeftCh, _ = EgressPipe.CreateEgressPipeLine(MgmtLeftCh)
	// Data from Outside world is fed here to this Ingress Channel. Data can
	// from Web server as HTTP Login Page Response
	//req := &Commondefs.Request{User: &Commondefs.UserInfo{"Raju", "PRaju"}}
	//
	time.Sleep(2 * time.Second)
	resp := &Commondefs.LoginResponse{}
	go SendData(resp, GlobalIngressCh, MgmtLeftCh)

	job := &Commondefs.UserJob{UserName: "Raju"}
	go SendData(job, GlobalIngressCh, MgmtLeftCh)

	time.Sleep(10 * time.Second)
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
