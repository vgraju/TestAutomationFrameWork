package EgressPipe

import (
	"fmt"
)

// Creates all the go routines for mgmt pipeline
func CreateEgressPipeLine(inputCh chan interface{}) (chan interface{}, error) {
	var leftCh chan interface{}

	fmt.Println("Inside Create Egress Pipeline")
	leftCh, _ = CreateDecisionMaking(inputCh)
	leftCh, _ = CreateFinalResult(inputCh)
	return leftCh, nil
}
func CreateDecisionMaking(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateDecisionMaking")
		for val := range inputCh {
			// Business Logic Here
			// Post the data onto left channel
			leftCh <- val
		}
	}()
	return leftCh, nil
}
func CreateFinalResult(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateFinalResult")
		for val := range inputCh {
			// Business Logic Here
			// Post the data onto left channel
			leftCh <- val
		}
	}()
	return leftCh, nil
}
