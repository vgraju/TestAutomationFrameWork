package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

type MySsh struct {
	client  *ssh.Client
	r       io.Reader
	w       io.Writer
	session *ssh.Session
	WriteCh chan string
	ReadCh  chan string
	log     *log.Logger
}

func main() {

	var wg sync.WaitGroup
	var MyIP = []string{"192.168.100.12", "192.168.100.18"}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {

			logPath := "/Users/rajuvegesana/vgraju/TestAutomationFrameWork/Logs/"
			logPath = logPath + MyIP[i] + ".txt"
			file, err := os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("COiult not create File ", logPath, "err:", err)
				return
			}

			defer file.Close()

			w := bufio.NewWriter(file)
			elem := MySsh{log: log.New(w, "logger", log.Lshortfile)}
			elem.log.SetOutput(file)

			fmt.Println("Log path ", logPath, "elem", elem)

			elem.ReadCh = make(chan string)
			elem.WriteCh = make(chan string)
			config := &ssh.ClientConfig{User: "admin", Auth: []ssh.AuthMethod{ssh.Password("mysnaproute")},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(), Timeout: 0}
			fmt.Println("Client:", config)

			client, err := ssh.Dial("tcp", MyIP[i]+":22", config)
			if err != nil {
				fmt.Println("Connedtion failed :", err)
			} else {

				fmt.Println("Connedtion Success:", err)
				elem.client = client
			}

			elem.CreateSession()

			output := elem.SendCommand("conf")
			fmt.Printf(output)

			for vlan := 80; vlan < 200; vlan++ {
				cmd := fmt.Sprintf("no vlan %d", vlan)
				elem.log.Println(cmd)

				output = elem.SendCommand(cmd)
				fmt.Printf("%s, i=%d", output, i)
			}
			output = elem.SendCommand("apply")
			fmt.Printf(output)
			wg.Done()
		}(i)
	}

	fmt.Println("waiting for go routines to exit")
	wg.Wait()
}

/*
func ReadUntil(sshOut io.Reader, expString string) string {
	scanner := bufio.NewScanner(sshOut)

	var buff bytes.Buffer
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Reading from STDOUT", text)
		if strings.Contains(text, expString) {
			break
		}
		buff.Write([]byte(text))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR", err)
	}

	return buff.String()
}
*/
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
func (c *MySsh) CreateSession() {
	session, _ := c.client.NewSession()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	sshIn, _ := session.StdinPipe()
	c.w = sshIn
	sshOut, _ := session.StdoutPipe()
	c.r = sshOut

	if err := session.RequestPty("xterm", 0, 200, modes); err != nil {
		fmt.Println("Failed1")
		session.Close()
		return
	}

	if err := session.Shell(); err != nil {
		fmt.Println("Failed2")
		session.Close()
		return
	}

	c.session = session

	c.MyGoRoutine()
	return
}
func (c *MySsh) SendCommand(cmd string) string {

	cmd = cmd + "\n"
	c.WriteCh <- cmd
	return strings.TrimSpace(<-c.ReadCh)
}

func (c *MySsh) MyGoRoutine() {
	go func() {
		for val := range c.WriteCh {

			c.w.Write([]byte(val))
			//time.Sleep(time.Second * 1)
			time.Sleep(time.Millisecond * 100)

			str := ReadUntil(c.r, "$")
			c.ReadCh <- str

		}
	}()
}
