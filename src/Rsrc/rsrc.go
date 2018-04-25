package Rsrc

import "fmt"

func (rsc *Resource) Add() {
	// THis will Add resources to the Pool
	//

}

func (rsc *Resoure) GetResoures(count int) ([]string, error) {
	return []string{}, nil
}

// This go routine, loops for ever and listens to Resources
// Adds During System Bringup (setupfile), When job is done
// resource is added to the pool, and when job is take,
// resrouce is taken out from the pool
func ResourceRoutine() {
	fmt.Println("Creating Resource Go routine")
	for {
		// Create a channel to read from the file for resoureces.
		// During init, we send message so that this will
		// read the configuration file and build the DUT Database.

		// Channel For Resoure Done API (this will post here)
		// Channel for Resoure Take ??

		// THis will post the message to Resoure Channels where
		// users are waiting for jobs - 1ResrcQ, 2 ResrcQ etc..
		// FOr Basic Sanity- we need 1 ResrcQ, L2 - we need 2 etc.
	}
}
