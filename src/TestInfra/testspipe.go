package TestInfra

import (
	"fmt"
)

//
// Creates all the go routines for mgmt pipeline
func CreateIngressTestPipeLine(inputCh chan interface{}) (chan interface{}, error) {
	var leftCh chan interface{}

	fmt.Println("Inside Create Ingreess Test Pipeline")
	leftCh, _ = CreateBasicSanityInput(inputCh)
	leftCh, _ = CreateL2Tests(inputCh)
	leftCh, _ = CreateL3Tests(inputCh)
	return leftCh, nil
}
func CreateBasicSanityInput(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateBasicSanityInput")
		for val := range inputCh {
			fmt.Println("data on: CreateBasicSanityInput")
			// Business Logic Here
			// Post the data onto left channel
			leftCh <- val
		}
	}()
	return leftCh, nil
}
func CreateL2Tests(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateL2Tests")
		for val := range inputCh {
			fmt.Println("data on: CreateL2Tests")
			// Business Logic Here
			// Post the data onto left channel
			leftCh <- val
		}
	}()
	return leftCh, nil
}
func CreateL3Tests(inputCh chan interface{}) (chan interface{}, error) {
	leftCh := make(chan interface{})
	go func() {
		fmt.Println("Go Routine : CreateL3Tests")
		for val := range inputCh {
			fmt.Println("data on: CreateL3Tests")
			// Business Logic Here
			// Post the data onto left channel
			leftCh <- val
		}
	}()
	return leftCh, nil
}
