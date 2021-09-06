package models

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func intiSky() {
	c := cron.New(cron.WithSeconds()) //精确到秒

	//定时任务
	spec := "0 0 8/12 * * ?" //cron表达式，每秒一次
	c.AddFunc(spec, func() {
		fmt.Println("开始wskey转换")

	})

	c.Start()
}
