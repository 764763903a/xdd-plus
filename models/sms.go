package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"regexp"
	"time"
)

var codes map[string]chan string

type Query struct {
	Ck struct {
		PtPin interface{} `json:"ptPin"`
		PtKey interface{} `json:"ptKey"`
		Empty bool        `json:"empty"`
	} `json:"ck"`
	PageStatus        string `json:"pageStatus"`
	AuthCodeCountDown int    `json:"authCodeCountDown"`
	CanClickLogin     bool   `json:"canClickLogin"`
	CanSendAuth       bool   `json:"canSendAuth"`
	SessionTimeOut    int    `json:"sessionTimeOut"`
	AvailChrome       int    `json:"availChrome"`
}

type Session struct {
	Value string
}

func (sess *Session) create() error {
	var url = "https://github.com/rubyangxg/jd-qinglong"
	if Config.SMSAddress == "" {
		return errors.New("未配置服务器地址，仓库地址：" + url)
	}
	html, _ := httplib.Get(Config.SMSAddress).String()
	res := regexp.MustCompile(`value="([\d\w]+)"`).FindStringSubmatch(html)
	if len(res) == 0 {
		return errors.New("崩了请找作者，仓库地址：https://github.com/rubyangxg/jd-qinglong")
	}
	sess.Value = res[1]
	return nil
}

func (sess *Session) control(name, value string) error {
	req := httplib.Post(Config.SMSAddress + "/control")
	req.Param("currId", name)
	req.Param("currValue", value)
	req.Param("clientSessionId", sess.String())
	_, err := req.String()
	// fmt.Println("controll", name, value, rt)
	return err
}

func (sess *Session) login(phone, sms_code string) error {
	req := httplib.Post(Config.SMSAddress + "/jdLogin")
	req.Param("phone", phone)
	req.Param("sms_code", sms_code)
	req.Param("clientSessionId", sess.String())
	_, err := req.String()
	// fmt.Println(phone, sms_code, rt)
	return err
}

func (sess *Session) sendAuthCode() error {
	req := httplib.Get(Config.SMSAddress + "/sendAuthCode?clientSessionId=" + sess.String())
	_, err := req.Response()
	return err
}

func (sess *Session) String() string {
	return sess.Value
}

func (sess *Session) query() (*Query, error) {
	query := &Query{}
	// fmt.Println(sess.String(), "+++")
	data, err := httplib.Get(fmt.Sprintf("%s/getScreen?clientSessionId=%s", Config.SMSAddress, sess.String())).Bytes()
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(data))
	err = json.Unmarshal(data, &query)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func (sess *Session) Phone(phone string) error {
	err := sess.create()
	if err != nil {
		return err
	}
	for {
		query, err := sess.query()
		if err != nil {
			return err
		}
		if query.PageStatus == "NORMAL" {
			break
		}
		if query.PageStatus == "SESSION_EXPIRED" {
			return sess.Phone(phone)
		}
		time.Sleep(time.Second)
	}
	err = sess.control("phone", phone)
	if err != nil {
		return err
	}
	return nil
}

func (sess *Session) SmsCode(sms_code string) error {
	err := sess.control("sms_code", sms_code)
	if err != nil {
		return err
	}
	return nil
}

func (sess *Session) crackCaptcha() error {
	_, err := httplib.Get(fmt.Sprintf("%s/crackCaptcha?clientSessionId=%s", Config.SMSAddress, sess.String())).Response()
	return err
}
