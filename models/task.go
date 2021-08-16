package models

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

type Task struct {
	ID      int
	Cron    string
	Path    string
	Enable  bool
	Mode    string //obo alo
	Word    string
	Run     func()
	Name    string
	Timeout int
	Envs    []Env
}

type Env struct {
	Name  string
	Value string
}

func initTask() {
	for i := range Config.Tasks {
		if Config.Tasks[i].Name == "" {
			slice := strings.Split(Config.Tasks[i].Path, "/")
			len := len(slice)
			if len == 0 {
				return
			}
			Config.Tasks[i].Name = slice[len-1]
		}
		// runTask(&Config.Tasks[i])
		createTask(&Config.Tasks[i])
	}
}

func createTask(task *Task) {
	id, err := c.AddFunc(task.Cron, func() {
		runTask(task)
	})
	if err != nil {
		logs.Warn(task.Word, "任务创建失败")
	} else {
		task.ID = int(id)
		logs.Info(task.Word, "任务创建成功")
	}

}

func runTask(task *Task) {
	var path = ExecPath + "/scripts/" + task.Name
	if strings.Contains(task.Path, "http") {
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			logs.Warn("打开%s失败，", path, err)
		}
		url := task.Path
		if strings.Contains(url, "raw.githubusercontent.com") {
			url = GhProxy + url
		}
		r, err := httplib.Get(url).Response()
		if err != nil {
			logs.Warn("下载%s失败，", task.Path, err)
		}
		io.Copy(f, r.Body)
		f.Close()
	}
	lan := Config.Node
	if strings.Contains(task.Name, ".py") {
		lan = Config.Python
	}
	envs := ""
	for _, env := range task.Envs {
		envs += fmt.Sprintf("export %s=%s", env.Name, env.Value)
	}
	sh := fmt.Sprintf(`
%s
cd %s
%s %s
	`, envs,
		ExecPath+"/scripts/",
		lan, task.Name)
	cmd := exec.Command("sh", "-c", sh)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("path: %s", cmd.Path)
	err := cmd.Run()
	if err != nil {
		logs.Warn("%v", err)
	}
}
