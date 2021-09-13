package models

import (
	"bufio"
	"io"
	"os/exec"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func cmd(str string, sender *Sender) string {
	cmd := exec.Command("sh", "-c", str)
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		logs.Warn("cmd.StdoutPipe: ", err)
		return err.Error()
	}
	cmd.Dir = ExecPath + "/scripts/"
	err = cmd.Start()
	if err != nil {
		logs.Warn("%v", err)
		return err.Error()
	}
	go func() {
		msg := ""
		reader := bufio.NewReader(stderr)
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			msg += line
		}
		if msg != "" {
			sender.Reply(msg)
		}
	}()
	msg := ""
	reader := bufio.NewReader(stdout)
	st := time.Now()
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		msg += line
		nt := time.Now()
		if (nt.Unix() - st.Unix()) > 15 {
			sender.Reply(msg)
			st = nt
			msg = ""
		}
	}
	if msg != "" {
		sender.Reply(msg)
	}
	err = cmd.Wait()
	return msg
}
