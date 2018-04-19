package Commondefs

type GORESP int

const (
	PASSTHROUGH_RESP GORESP = iota
	USER_RESP
)

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
