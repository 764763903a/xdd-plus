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

func (rp *Repo) init() {
	rp.filename = strings.Replace(strings.Replace(rp.Git, "https://", "", -1), "/", "_", -1)
}

func initRepos() {
	for _, repo := range Config.Repos {
		repo.init()
		if _, err := os.Stat(ExecPath + "/" + repo.filename); err != nil {
			repo.gitClone()
		}
	}
}

func (rp *Repo) gitClone() {
	cmd := exec.Command("git", "clone", rp.Git, rp.filename)
	cmd.Path = ExecPath + "/repos"
	fmt.Println(cmd.Path)
	fmt.Println(cmd.Start())
}
