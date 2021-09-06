package models

import (
	"log"

	"github.com/robfig/cron/v3"
)

func intiSky() {
	log.Println("Starting...")
	c := cron.New() // 新建一个定时任务对象
	c.AddFunc("0 59 0/12 * * *", func() {
		log.Println("hello world")
	}) // 给对象增加定时任务
	c.Start()
}
