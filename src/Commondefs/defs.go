package Commondefs

type GORESP int

const (
	PASSTHROUGH_RESP GORESP = iota
	USER_RESP
)

// Global workeri Buffered channels  on which users
// will be waiting on
type WorkerChannels struct {
	RsrcChan1 chan interface{} // All these users needs 1 rsrc - exm: Basic Sanity
	RsrcChan2 chan interface{} // Users needs 2 rsrcs - exm L2 etc
	RsrcChan3 chan interface{}
}
type UserJob struct {
	// User Information Needed for Job
	UserName string
	// Need what test that user have selected
	// example : L2 - True, L3-False,BasicSanity-True
	TestCases map[string]struct{}
}

type UserInfo struct {
	UName string
	UPwd  string
}

type TestModRequirement struct {
	TestMod map[string]int // Number of Dut's need for TESTMOD- Say BSANITY:2 etc
}
type Request struct {
	User           *UserInfo
	ResoureSlice   []string            // Resources that this Req is currently using - A,B..
	TestModSlice   []string            // Test cases that this Req need to go through, BSANITY, L2 ..
	TestModDoneMap map[string]struct{} // once the Testcase is finished, this will be set to Done
	ChartInfo      string              // Chart information that have to be pulled after getting rsrc
	ExeType        int                 // Const - Use CLI or Rest or Kubectl for configuraation
	FailedTests    []string            // All the failed cases are added here. Each Testcase is Go routine ??
	RetryCount     int                 // On Failures, Try many retries, say:2, Try 2 times, if failed case
}

type LoginResponse struct {
	User *UserInfo
}

// Login Response from backend to the user. THis should display the
// next page - with Drop box/list of Testcase that user can select etc..
type HttpUserResponse struct {
}
