package Mgmt

import (
	"Commondefs"
	"fmt"
)

// Creates all the go routines for mgmt pipeline
func CreateManagementPipeLine(inputCh chan interface{}) (chan interface{}, error) {
	var leftCh chan interface{}

	leftCh, _ = CreateIngressInput(inputCh)
	leftCh, _ = CreateUserMgmt(leftCh)
	leftCh, _ = CreateRequestRateLimiter(leftCh)
	leftCh, _ = CreateResoureManager(leftCh)
	fmt.Println("Inside Create Management Pipeline")
	return leftCh, nil
}

func CreateIngressInput(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateIngressInput")
		for val := range inputCh {
			// Business Logic Here
			// Post the data onto left channel
			leftCh <- val
		}
	}()
	return leftCh, nil
}
func CreateUserMgmt(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateUserMgmt")
		for val := range inputCh {
			// Check whether the data is PASSTHROUGH
			// based on response. If this data is for my block
			// in Pipeline, then we will modify the data
			val, resp := ProcessUserMgmt(val)
			if resp == Commondefs.USER_RESP {
				fmt.Println("Proessing User Response in this Routine")

				// SEnd the response back to the right side user
			} else {
				// Post the data onto left channel
				leftCh <- val
			}
		}
	}()
	return leftCh, nil
}
func CreateRequestRateLimiter(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateRequestRateLimiter")
		for val := range inputCh {
			// Business Logic Here
			// Post the data onto left channel
			leftCh <- val
		}
	}()
	return leftCh, nil
}
func CreateResoureManager(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateResoureManager")
		for val := range inputCh {
			// Business Logic Here
			// Post the data onto left channel
			leftCh <- val
		}
	}()
	return leftCh, nil
}

func ProcessUserMgmt(data interface{}) (interface{}, Commondefs.GORESP) {

	switch data.(type) {
	case *Commondefs.LoginResponse:
		fmt.Println("Processing Login Response Recevied")
		// Check the User Exists in database
		// Send the request to TMPL to get resource selection page
		// and send back to user and not to the leftCh
		// Send Template here
		return &Commondefs.RequirementResponse{}, Commondefs.USER_RESP
	case *Commondefs.RequirementResponse:
		fmt.Println("Processing RequirementResponse Recevied")
	default:
		fmt.Println("Data not changed")
	}
	return data, Commondefs.PASSTHROUGH_RESP
}
