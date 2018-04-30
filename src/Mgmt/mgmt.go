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
	leftCh, _ = IngressAuditPipe(leftCh)
	fmt.Println("Inside Create Management Pipeline")
	return leftCh, nil
}

func CreateIngressInput(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateIngressInput")
		for val := range inputCh {
			fmt.Println("data on CreateIngressInput")
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
			fmt.Println("data on CreateUserInput")
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
			fmt.Println("data on CreateRequestRateLimiter")
			// Business Logic Here
			// Post the data onto left channel
			leftCh <- val
		}
	}()
	return leftCh, nil
}
func IngressAuditPipe(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : IngressAuditPipe")
		for val := range inputCh {
			fmt.Println("data on Ingress Audit Pipe")
			// Business Logic Here
			// Post the data onto left channel
			RecordIngressLogs(val)
			leftCh <- val
		}
	}()
	return leftCh, nil
}

func RecordIngressLogs(data interface{}) (interface{}, Commondefs.GORESP) {
	fmt.Println("Processing Ingress Logs")
	// Add Switch / Remove Switch
	// Keep a Global Slice of Resources and Remove the required
	// Number of resource from this slice and

	return data, Commondefs.PASSTHROUGH_RESP
}
func ProcessUserMgmt(data interface{}) (interface{}, Commondefs.GORESP) {

	switch data.(type) {
	case *Commondefs.LoginResponse:
		fmt.Println("Processing Login Response Recevied")
		// Check the User Exists in database
		// Send the request to TMPL to get resource selection page
		// and send back to user and not to the leftCh
		// Send Template here
		return &Commondefs.HttpUserResponse{}, Commondefs.USER_RESP
	case *Commondefs.HttpUserResponse:
		fmt.Println("Processing RequirementResponse Recevied")
	default:
		fmt.Println("Data not changed")
	}
	return data, Commondefs.PASSTHROUGH_RESP
}
