package models

import (
	"bufio"
	"io"
	"os/exec"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func cmd(str string, msgs ...interface{}) string {
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
			sendMessagee(msg, msgs...)
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
		if (nt.Unix() - st.Unix()) > 1 {
			go sendMessagee(msg, msgs...)
			st = nt
			msg = ""
		}
	}
	if msg != "" {
		sendMessagee(msg, msgs...)
	}
	err = cmd.Wait()
	return msg
}
