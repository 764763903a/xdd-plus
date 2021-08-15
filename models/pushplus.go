package models

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/httplib"
)

func pushPlus(token string, content string) {
	req := httplib.Post("http://pushplus.hxtrip.com/send")
	req.Header("Content-Type", "application/json")
	data, _ := json.Marshal(struct {
		Token   string `json:"token"`
		Content string `json:"Content"`
	}{
		Token:   token,
		Content: content,
	})
	req.Body(data)
	req.Response()
}
