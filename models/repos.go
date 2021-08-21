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
}

var reposPath = ""

func (rp *Repo) init() {
	rp.filename = strings.Replace(strings.Replace(rp.Git, "https://", "", -1), "/", "_", -1)
}

func initRepos() {
	reposPath = ExecPath + "/repos"
	if _, err := os.Stat(reposPath); err != nil {
		os.MkdirAll(reposPath, os.ModePerm)
	}
	for _, repo := range Config.Repos {
		repo.init()
		if _, err := os.Stat(reposPath + "/" + repo.filename); err != nil {
			repo.gitClone()
		}
	}
}

func (rp *Repo) gitClone() {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("git clone %s %s", rp.Git, rp.filename))
	cmd.Output()
}
