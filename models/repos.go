package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/robfig/cron/v3"
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
			if strings.Contains(Config.Repos[i].gitPull(), "changed") {
				Config.Repos[i].addTask()
			}
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
	cmd(fmt.Sprintf("cd %s && git clone %s %s", reposPath, rp.Git, rp.filename), &Sender{})
}

func (rp *Repo) gitPull() string {
	return cmd(fmt.Sprintf("cd %s && git stash && git pull", rp.path), &Sender{})
}

func (rp *Repo) cpConfig() {
	for _, js := range []string{"jdCookie", "jdFruitShareCodes", "jdPetShareCodes", "jdPlantBeanShareCodes", "jdFactoryShareCodes", "jdDreamFactoryShareCodes", "jdJxncShareCodes"} {
		cmd(fmt.Sprintf(`cp `+js+`.js %s`, rp.path+"/"+js+".js"), &Sender{})
	}
}

func (rp *Repo) addTask() {
	dir_list, e := ioutil.ReadDir(rp.path)
	if e != nil {
		return
	}
	nts := []Task{}
	for _, v := range dir_list {
		if strings.Contains(v.Name(), ".js") {
			f, err := os.Open(rp.path + "/" + v.Name())
			if err != nil {
				continue
			}
			data, _ := ioutil.ReadAll(f)
			f.Close()
			res := regexp.MustCompile(`([\d\-,\*]+ [\d\-,\*]+ [\d\-,\*]+ [*]+ [*]+)[\s\S]+Env[(]['"]([^'"]+)['"][)]`).FindStringSubmatch(string(data))
			if len(res) > 0 {
				nts = append(nts, Task{
					Cron:  res[1],
					Name:  v.Name(),
					Title: res[2],
					Git:   rp.path,
				})
			}
		}
	}
	for i := range rp.Task {
		if rp.Task[i].ID != 0 {
			c.Remove(cron.EntryID(rp.Task[i].ID))
		}
	}
	rp.Task = nts
	for i := range rp.Task {
		task := &rp.Task[i]
		eid, err := c.AddFunc(task.Cron, func() {
			// if Cdle {
			// 	return
			// }
			logs.Info("执行任务 %s %s ", task.Title, task.Cron)
			runTask(task, &Sender{})
		})
		if err == nil {
			logs.Info("添加任务 %s %s ", rp.Task[i].Title, rp.Task[i].Cron)
			rp.Task[i].ID = int(eid)
		}
	}
}
