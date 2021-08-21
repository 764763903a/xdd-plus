package models

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/robfig/cron/v3"
)

var c *cron.Cron

func initCron() {
	c = cron.New()
	if Config.DailyAssetPushCron != "" {
		_, err := c.AddFunc(Config.DailyAssetPushCron, DailyAssetsPush)
		if err != nil {
			logs.Warn("资产推送任务失败：%v", err)
		} else {
			logs.Info("资产推送任务就绪")
		}
		c.AddFunc("3 */1 * * *", initVersion)
		c.AddFunc("40 */1 * * *", GitPullAll)

	}
	c.Start()
}
