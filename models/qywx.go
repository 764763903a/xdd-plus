package models

import (
	"time"

	"github.com/beego/beego/v2/client/httplib"
)

type QywxConfig struct {
	QywxKey string
	Content string
}

type QywxNotifyMessage struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func qywxNotify(c *QywxConfig) {
	if c.QywxKey == "" {
		return
	}
	wx := QywxNotifyMessage{
		Msgtype: "text",
	}
	wx.Text.Content = c.Content
	req := httplib.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + c.QywxKey)
	req.Header("Content-Type", "application/json")
	req, _ = req.JSONBody(wx)
	req.SetTimeout(time.Second*2, time.Second*2)
	req.Response()
}
