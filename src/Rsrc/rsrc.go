package Rsrc

import (
	ssh "SshClient"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Need Multiple Queue - 1DutQ, 2DUTQ etc...and users
// are added to these queue right away...(buffered Channels)
// These are all waiting on 1QRsrcChan, 2QRsrcChan etc...whic
// will be notified from Resoures availabilty - Weighted
// Round robin. Along with this one High Priority Queue,
// when Admin User is using this Testbed for Sanity.

func Init() *ResourceMgmt {
	// This API create a new Resource, where any package
	// wants to use this, will need this resource

	// Loop ever Looptime seconds looking for resrouces
	rscInfo := ResourceInfo{rsrcInfo: make(map[string]string)}
	return &ResourceMgmt{FileReadCh: make(chan struct{}),
		RsrcAddCh: make(chan *ResourceInput), RscMap: rscInfo,
		LoopTime: 5} // Loop ever 5 seconds for searching
}

// If this Package needs any references from other packages,
// then the main will call with all the refercnces and this Pakcage
// will store trhe required obj referece of remote pkgs
func (rsrc *ResourceMgmt) SetRefs(object []interface{}) {
	for _, value := range object {
		switch value.(type) {
		case *ssh.SshClientInfo:
			fmt.Println("In Resoure Mgmemt for sshclient referece")
			rsrc.sshObj = value.(*ssh.SshClientInfo)
		}
	}
}
func (rsc *ResourceMgmt) Add() {
	// THis will Add resources to the Pool
	//

}

func (rsc *ResourceMgmt) GetResoures(count int) ([]string, error) {
	return []string{}, nil
}

// This go routine, loops for ever and listens to Resources
// Adds During System Bringup (setupfile), When job is done
// resource is added to the pool, and when job is take,
// resrouce is taken out from the pool

func (rsc *ResourceMgmt) ResourceRoutine() {
	fmt.Println("Creating Resource Go routine")
	for {
		// Create a channel to read from the file for resoureces.
		// During init, we send message so that this will
		// read the configuration file and build the DUT Database.
		SetupFile := "/Users/rajuvegesana/vgraju/TestAutomationFrameWork/setup.txt"
		select {
		case <-rsc.FileReadCh:
			fmt.Println("BUding Resource Map by Reading file:")
			file, err := os.OpenFile(SetupFile, os.O_RDONLY, 0666)
			if err != nil {
				fmt.Println("Error in Reading File. Err:", err)
				return
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				text := strings.TrimSpace(scanner.Text())
				//fmt.Println("reading file", text)
				// If there are no comments then process the text
				if !strings.Contains(text, "#") {
					if strings.Contains(text, "IPADDR") {
						rsc.CreateRsrcFromFile(text)
						// Increment resource count
						rsc.RsrcCount = rsc.RsrcCount + 1

					}

				}
			}

		case resource := <-rsc.RsrcAddCh:
			fmt.Println("Resource Add ", resource)
			// Check whether the resource already exists, else don't
			// increment
			rsc.RsrcCount = rsc.RsrcCount + 1
		case <-time.After(rsc.LoopTime * time.Second):
			fmt.Println("Looking for Jobs ")
		}
		// Channel For Resoure Done API (this will post here)
		// Channel for Resoure Take ??

		// THis will post the message to Resoure Channels where
		// users are waiting for jobs - 1ResrcQ, 2 ResrcQ etc..
		// FOr Basic Sanity- we need 1 ResrcQ, L2 - we need 2 etc.

	}
}

func (rsrc *ResourceMgmt) CreateRsrcFromFile(str string) {
	//NAME:DUTA,IPADDR:192.168.100.18
	strSlice := strings.Split(str, ",")
	for _, value := range strSlice {
		// IPADDR:192.168.100.18
		element := strings.Split(value, ":")
		rsrc.RscMap.rsrcInfo[element[0]] = element[1]
		// Name : Add to ResourceSlice
		if strings.ToLower(element[0]) == "name" {
			rsrc.RscSlice = append(rsrc.RscSlice, element[1])
			fmt.Println("Resource Slice ", rsrc.RscSlice)
		}

	}
	rsrc.sshObj.CreateClients(rsrc.RscMap.rsrcInfo)
	fmt.Println("Rsrcmap ", rsrc.RscMap)
}
