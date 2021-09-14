package models

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/adapter/logs"
)

func Daemon() {
	args := os.Args[1:]
	execArgs := make([]string, 0)
	l := len(args)
	for i := 0; i < l; i++ {
		if strings.Index(args[i], "-d") == 0 {
			continue
		}

		execArgs = append(execArgs, args[i])
	}
	proc := exec.Command(os.Args[0], execArgs...)
	err := proc.Start()
	if err != nil {
		panic(err)
	}
	logs.Info("小滴滴运行于后台模式")
	os.Exit(0)
}

func killp() {
	pids, err := ppid()
	if err == nil {
		if len(pids) == 0 {
			return
		} else {
			exec.Command("sh", "-c", "kill -9 "+strings.Join(pids, " ")).Output()
		}
	} else {
		return
	}
}

func ppid() ([]string, error) {
	pid := fmt.Sprint(os.Getpid())
	pids := []string{}
	rtn, err := exec.Command("sh", "-c", "pidof "+pname).Output()
	if err != nil {
		return pids, err
	}
	re := regexp.MustCompile(`[\d]+`)
	for _, v := range re.FindAll(rtn, -1) {
		if string(v) != pid {
			pids = append(pids, string(v))
		}
	}
	return pids, nil
}
