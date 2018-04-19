package main

import (
	"Commondefs"
	"Mgmt"
	"fmt"
)

func main() {

	// Create Ingress Channel
	GlobalIngressCh := make(chan interface{})
	// Create Mgmt Pipeline - User managerment, Resource Management etc
	MgmtLeftCh, _ := Mgmt.CreateManagementPipeLine(GlobalIngressCh)
	// Data from Outside world is fed here to this Ingress Channel. Data can
	// from Web server as HTTP Login Page Response
	//req := &Commondefs.Request{User: &Commondefs.UserInfo{"Raju", "PRaju"}}
	resp := &Commondefs.LoginResponse{}
	GlobalIngressCh <- resp
	recvData := <-MgmtLeftCh

	DecodeReceivedData(recvData)
	// wait for ever
	//	select {}
}

func DecodeReceivedData(data interface{}) {
	switch data.(type) {
	case *Commondefs.Request:
		fmt.Println("Request Recevied")
	case *Commondefs.LoginResponse:
		fmt.Println("Login Reponse Recevied")
	case *Commondefs.RequirementResponse:
		fmt.Println("RequirementResponse Recevied")
	default:
		fmt.Println("UNDEFINED RESPONSE RECEIVED:")
	}
	return
}
