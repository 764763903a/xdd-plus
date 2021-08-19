package models

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/httplib"
)

func pushPlus(token string, content string) {
	if token == "" {
		return
	}
	data, _ := json.Marshal(struct {
		Token    string `json:"token"`
		Content  string `json:"content"`
		Template string `json:"template"`
	}{
		Token:    token,
		Content:  content,
		Template: "txt",
	})
	req := httplib.Post("http://pushplus.hxtrip.com/send")
	req.Header("Content-Type", "application/json")
	req.Body(data)
	req.Response()
	req = httplib.Post("http://www.pushplus.plus/send")
	req.Header("Content-Type", "application/json")
	req.Body(data)
	req.Response()
}
