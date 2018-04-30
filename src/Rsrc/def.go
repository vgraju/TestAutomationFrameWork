package Rsrc

import (
	"SshClient"
	"time"
)

const (
	RESOURCEADD int = iota
	RESOURCEDEL
)

// Information of each resource, like
// IP addres, Console Addr, Rsc Type - Cisco/Flex etc
type ResourceInfo struct {
	rsrcInfo map[string]string
}
type ResourceMgmt struct {
	FileReadCh chan struct{}
	RsrcAddCh  chan *ResourceInput // We add resource over this channel
	RsrcCount  int                 // Total Resources currently in the system
	RscSlice   []string            // Slice of DUT's
	RscMap     ResourceInfo
	RscPortMap map[string][]string // "AB":"fpPort1, fpPort2..We also store BA etc.."
	sshObj     *SshClient.SshClientInfo
	LoopTime   time.Duration // After this time, we look for any jobs..in loop
}

// To Add/Del any New resource, we use this struct
type ResourceInput struct {
	Oper       int           // Resource Add/ Resource Del
	RscAInfo   *ResourceInfo // Name Of ResourceA
	RscBName   *ResourceInfo // Resource B optional
	RscPortMap []string      // fpPort1 To fpPort2 or fpPort1-fpPort5 To fpPort3-fpPort6. To is Must
}
