package SshClient

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func ReadUntil(sshOut io.Reader, expString string) string {
	buf := make([]byte, 1000)
	readStr := ""
	for {
		n, err := sshOut.Read(buf)
		readStr = readStr + string(buf[:n])
		if err == io.EOF || n == 0 {
			fmt.Println("Break,", err, n)
			break
		}

		if strings.Contains(readStr, expString) {
			break
		}

	}
	return readStr
}

func (c *ClientInfo) SendCommand(cmd string) string {

	cmd = cmd + "\n"
	c.WriteCh <- cmd
	return strings.TrimSpace(<-c.ReadCh)
}

func (c *ClientInfo) SessionGoRoutine() {
	go func() {
		fmt.Println("Starting Subprocess Go routine")
		for val := range c.WriteCh {

			c.w.Write([]byte(val))
			time.Sleep(time.Millisecond * 100)

			str := ReadUntil(c.r, "$")
			c.ReadCh <- str

		}
		fmt.Println("Ending Subprocess Go routine")
	}()
}
