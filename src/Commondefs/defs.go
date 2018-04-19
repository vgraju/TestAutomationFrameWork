package Commondefs

type GORESP int

const (
	PASSTHROUGH_RESP GORESP = iota
	USER_RESP
)

type UserJob struct {
	// User Information Needed for Job
	UserName string
	// Need what test that user have selected
	// example : L2 - True, L3-False,BasicSanity-True
	TestCases map[string]struct{}
}
type Resource struct {
}
type UserInfo struct {
	UName string
	UPwd  string
}
type Request struct {
	User *UserInfo
}

type LoginResponse struct {
	User *UserInfo
}
type RequirementResponse struct {
	// L2, L3, BSANITY, ALL
	TestModules map[string]struct{}
}
