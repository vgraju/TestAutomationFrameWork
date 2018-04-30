package Tests

// Create dynamically a Go Routine for each Testcase
func L2Tests() {
	// This is one time Cfg Test Slice. This can be update by Using Mutex
	var L2TestSlice = []string{"VerifyVlanCreate", "VerifyVlanDelete"}
	l2TestCount := len(L2TestSlice)
	for index := 0; index < l2TestCount; index++ {

	}
}
