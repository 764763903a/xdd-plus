package models

import (
	"fmt"
	"os"
	"strings"
)

type Repo struct {
	Git      string
	filename string
	Task     []Task
	path     string
}

var reposPath = ""

func (rp *Repo) init() {
	rp.filename = strings.Replace(strings.Replace(strings.Replace(rp.Git, "https://", "", -1), "/", "_", -1), ".git", "", -1)
	if !rp.exist() {
		rp.gitClone()
	} else {
		rp.gitPull()
	}
	rp.path = reposPath + "/" + rp.filename
	CopyConfigAll()
	rp.addTask()
}

func (rp *Repo) exist() bool {
	if _, err := os.Stat(rp.path); err != nil {
		return false
	}
	return true
}

func initRepos() {
	reposPath = ExecPath + "/repos"
	if _, err := os.Stat(reposPath); err != nil {
		os.MkdirAll(reposPath, os.ModePerm)
	}
	for i := range Config.Repos {
		Config.Repos[i].init()
	}
}

func GitPullAll() {
	for i := range Config.Repos {
		if Config.Repos[i].exist() {
			Config.Repos[i].gitPull()
			Config.Repos[i].cpConfig()
		}
	}
}

func CopyConfigAll() {
	for i := range Config.Repos {
		if Config.Repos[i].exist() {
			Config.Repos[i].cpConfig()
		}
	}
}

func (rp *Repo) gitClone() {
	cmd(fmt.Sprintf("cd %s && git clone %s %s", reposPath, rp.Git, rp.filename))
}

func (rp *Repo) gitPull() {
	cmd(fmt.Sprintf("cd %s && git stash && git pull", rp.path))
}

func (rp *Repo) cpConfig() {
	cmd(fmt.Sprintf(`cp jdCookie.js %s`, rp.path+"/jdCookie.js"))
}

func (rp *Repo) addTask() {
	// dir_list, e := ioutil.ReadDir(rp.path)
	// if e != nil {
	// 	return
	// }
	// for _, v := range dir_list {
	// 	if strings.Contains(v.Name(), ".js") {
	// 		f, err := os.Open(rp.path + "/" + v.Name())
	// 		if err != nil {
	// 			continue
	// 		}
	// 		data, _ := ioutil.ReadAll(f)
	// 		f.Close()
	// 		// fmt.Println(data)
	// 		res := regexp.MustCompile("().*Env[(]['\"](\\S+)['\"][)]").FindStringSubmatch(string(data))
	// 		if len(res) > 0 {
	// 			fmt.Println(res[1])
	// 		}
	// 	}
	// }
}
