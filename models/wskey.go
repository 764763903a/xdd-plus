package models

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func intiSky() {
	c := cron.New(cron.WithSeconds()) //精确到秒

	//定时任务
	spec := "1 * * * * ?" //cron表达式，每秒一次
	c.AddFunc(spec, func() {
		fmt.Println("11111")
	})

	c.Start()
}
