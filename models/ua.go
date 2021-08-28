package models

import (
	"strings"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/httplib"
)

var ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 SP-engine/2.14.0 main%2F1.0 baiduboxapp/11.18.0.16 (Baidu; P2 13.3.1) NABar/0.0"

func initUserAgent() {
	u := &UserAgent{}
	err := db.Order("id desc").First(u).Error
	if err != nil && strings.Contains(err.Error(), "converting") {
		db.Migrator().DropTable(&UserAgent{})
		Daemon()
	}
	if u.Content != "" {
		ua = u.Content
	} else {
		if Config.UserAgent != "" {
			logs.Info("使用自定义User-Agent")
			ua = Config.UserAgent
		} else {
			logs.Info("更新User-Agent")
			var err error
			ua, err = httplib.Get(GhProxy + "https://raw.githubusercontent.com/cdle/xdd/main/ua.txt").String()
			if err != nil {
				logs.Info("更新User-Agent失败")
			}
		}
	}
}

func GetUserAgent() string {
	return ua
}

type UserAgent struct {
	ID      int
	Content string
}
