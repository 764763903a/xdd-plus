package models

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Repo struct {
	Git      string
	filename string
	ok       bool
}

var reposPath = ""

func (rp *Repo) init() {
	rp.filename = strings.Replace(strings.Replace(rp.Git, "https://", "", -1), "/", "_", -1)
	if !rp.exist() {
		rp.gitClone()
	}
}

func (rp *Repo) exist() bool {
	if _, err := os.Stat(reposPath + "/" + rp.filename); err != nil {
		return false
	}
	return true
}

func initRepos() {
	reposPath = ExecPath + "/repos"
	if _, err := os.Stat(reposPath); err != nil {
		os.MkdirAll(reposPath, os.ModePerm)
	}
	for _, repo := range Config.Repos {
		repo.init()
	}
}

func (rp *Repo) gitClone() {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("git clone %s %s", rp.Git, rp.filename))
	cmd.Path = reposPath
	fmt.Println("sh", "-c", fmt.Sprintf("git clone %s %s", rp.Git, rp.filename))
	fmt.Println(cmd.Output())
}
