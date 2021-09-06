package models

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/robfig/cron/v3"
)

func intiSky() {
	logs.Info("Starting...")
	c := cron.New() // 新建一个定时任务对象
	c.AddFunc("0 3 1 * * *", func() {
		logs.Info("hello world")
	}) // 给对象增加定时任务
	c.Start()
}
