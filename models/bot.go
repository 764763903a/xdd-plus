package models

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

var SendQQ = func(a int64, b interface{}) {

}
var SendQQGroup = func(aint64, b int64, c interface{}) {

}
var ListenQQPrivateMessage = func(uid int64, msg string) {
	SendQQ(uid, handleMessage(msg, "qq", int(uid)))
}

var ListenQQGroupMessage = func(gid int64, uid int64, msg string) {
	if gid == Config.QQGroupID {
		if Config.QbotPublicMode {

			SendQQGroup(gid, uid, handleMessage(msg, "qqg", int(gid), int(uid)))
		} else {
			SendQQ(uid, handleMessage(msg, "qq", int(uid)))
		}
	}
}

var replies = map[string]string{}

func InitReplies() {
	f, err := os.Open(ExecPath + "/conf/reply.php")
	if err == nil {
		defer f.Close()
		data, _ := ioutil.ReadAll(f)
		ss := regexp.MustCompile("`([^`]+)`\\s*=>\\s*`([^`]+)`").FindAllStringSubmatch(string(data), -1)
		for _, s := range ss {
			replies[s[1]] = s[2]
		}
	}
	if _, ok := replies["壁纸"]; !ok {
		replies["壁纸"] = "https://acg.toubiec.cn/random.php"
	}
}

var sendMessagee = func(msg string, msgs ...interface{}) {
	if len(msgs) == 0 {
		return
	}
	tp := msgs[1].(string)
	id := msgs[2].(int)
	switch tp {
	case "tg":
		SendTgMsg(id, msg)
	case "qq":
		SendQQ(int64(id), msg)
	case "qqg":
		SendQQGroup(int64(id), int64(msgs[3].(int)), msg)
	}
}

var sendAdminMessagee = func(msg string, msgs ...interface{}) {
	if len(msgs) == 0 {
		return
	}
	tp := msgs[1].(string)
	id := msgs[2].(int)
	switch tp {
	case "tg":
		if Config.TelegramUserID == id {
			SendTgMsg(id, msg)
		}
	case "qq":
		if int(Config.QQID) == id {
			SendQQ(int64(id), msg)
		}
	case "qqg":
		uid := msgs[3].(int)
		if int(Config.QQID) == uid {
			SendQQGroup(int64(id), int64(uid), msg)
		}
	}
}

var isAdmin = func(msgs ...interface{}) bool {
	if len(msgs) == 0 {
		return false
	}
	tp := msgs[1].(string)
	id := msgs[2].(int)
	switch tp {
	case "tg":
		if Config.TelegramUserID == id {
			return true
		}
	case "qq":
		if int(Config.QQID) == id {
			return true
		}
	case "qqg":
		uid := msgs[3].(int)
		if int(Config.QQID) == uid {
			return true
		}
	}
	return false
}

var handleMessage = func(msgs ...interface{}) interface{} {
	msg := msgs[0].(string)
	tp := msgs[1].(string)
	id := msgs[2].(int)
	switch msg {
	case "status", "状态":
		if !isAdmin(msgs...) {
			return "你没有权限操作"
		}
		return Count()
	case "sign", "打卡":
		return "打卡成功"
	case "qrcode", "扫码", "二维码", "scan":
		url := ""
		if tp == "qqg" {
			url = fmt.Sprintf("http://127.0.0.1:%d/api/login/qrcode.png?%vid=%v&qqguid=%v", web.BConfig.Listen.HTTPPort, tp, id, msgs[3].(int))
		} else {
			url = fmt.Sprintf("http://127.0.0.1:%d/api/login/qrcode.png?%vid=%v", web.BConfig.Listen.HTTPPort, tp, id)
		}
		rsp, err := httplib.Get(url).Response()
		if err != nil {
			return nil
		}
		return rsp
	case "升级", "更新", "update", "upgrade":
		if !isAdmin(msgs...) { //
			return "你没有权限操作"
		}
		if err := Update(msgs...); err != nil {
			return err.Error()
		}
		fallthrough
	case "重启", "reload", "restart":
		if !isAdmin(msgs...) {
			return "你没有权限操作"
		}
		sendAdminMessagee("小滴滴重启程序", msgs...)
		Daemon()
		return nil
	case "查询", "query":
		cks := GetJdCookies()
		tmp := []JdCookie{}
		for _, ck := range cks {
			if tp == "qq" {
				if ck.QQ == id {
					tmp = append(tmp, ck)
				}
			} else if tp == "qqg" {
				if ck.QQ == msgs[3].(int) {
					tmp = append(tmp, ck)
				}
			}
		}
		if len(tmp) == 0 {
			return "你尚未绑定🐶东账号，请对我说扫码，扫码后即可查询账户资产信息。"
		}
		for _, ck := range tmp {
			go sendMessagee(ck.Query(), msgs...)
		}
		return nil
	default:
		{ //tyt
			ss := regexp.MustCompile(`packetId=(\S+!!)`).FindStringSubmatch(msg)
			if len(ss) > 0 {
				runTask(&Task{Path: "jd_tyt.js", Envs: []Env{
					{Name: "tytpacketId", Value: ss[1]},
				}}, msgs...)
				return nil
			}
		}
		{ //
			ss := regexp.MustCompile(`pt_key=([^;=\s]+);pt_pin=([^;=\s]+)`).FindAllStringSubmatch(msg, -1)
			if len(ss) > 0 {
				xyb := 0
				for _, s := range ss {
					ck := JdCookie{
						PtKey: s[1],
						PtPin: s[2],
					}
					if CookieOK(&ck) {
						xyb++
						if tp == "qq" {
							ck.QQ = id

						} else if tp == "tg" {
							ck.Telegram = id
						} else if tp == "qqg" {
							ck.QQ = msgs[3].(int)
						}
						if nck, err := GetJdCookie(ck.PtPin); err == nil {
							nck.InPool(ck.PtKey)
							msg := fmt.Sprintf("更新账号，%s", ck.PtPin)
							(&JdCookie{}).Push(msg)
							sendMessagee("许愿币+1", msgs...)
							logs.Info(msg)
						} else {
							if Cdle {
								ck.Hack = True
							}
							NewJdCookie(&ck)
							msg := fmt.Sprintf("添加账号，%s", ck.PtPin)
							(&JdCookie{}).Push(msg)
							sendMessagee("许愿币+1", msgs...)
							logs.Info(msg)
						}
					}
				}
				go func() {
					Save <- &JdCookie{}
				}()
				return nil
			}
		}
		{
			s := regexp.MustCompile(`([^\s]+)\s+(.*)`).FindStringSubmatch(msg)
			if len(s) > 0 {
				v := s[2]
				switch s[1] {
				case "查询", "query":
					if !isAdmin(msgs...) {
						return "你没有权限操作"
					}
					cks := GetJdCookies()
					a := s[2]
					tmp := []JdCookie{}
					if s := strings.Split(a, "-"); len(s) == 2 {
						for i, ck := range cks {
							if i+1 >= Int(s[0]) && i+1 <= Int(s[1]) {
								tmp = append(tmp, ck)
							}
						}
					} else if x := regexp.MustCompile(`^[\s\d,]+$`).FindString(a); x != "" {
						xx := regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(a, -1)
						for i, ck := range cks {
							for _, x := range xx {
								if fmt.Sprint(i+1) == x[1] {
									tmp = append(tmp, ck)
								}
							}

						}
					} else {
						a = strings.Replace(a, " ", "", -1)
						for _, ck := range cks {
							if strings.Contains(ck.Note, a) || strings.Contains(ck.Nickname, a) || strings.Contains(ck.PtPin, a) {
								tmp = append(tmp, ck)
							}
						}
					}

					if len(tmp) == 0 {
						return "找不到匹配的账号"
					}
					for _, ck := range tmp {
						go sendMessagee(ck.Query(), msgs...)
					}
					return nil

				case "许愿":
					if tp == "qqg" {
						id = msgs[3].(int)
					}
					b := 0
					for _, ck := range GetJdCookies() {
						if id == ck.QQ || id == ck.Telegram {
							b++
						}
					}
					if b <= 0 {
						return "许愿币不足"
					} else {
						(&JdCookie{}).Push(fmt.Sprintf("%d许愿%s，许愿币余额%d。", id, v, b))
						return "收到许愿"
					}
				case "扣除许愿币":
					id, _ := strconv.Atoi(v)
					b := 0
					k := 0
					for _, ck := range GetJdCookies() {
						if id == ck.QQ || id == ck.Telegram {
							if k <= 5 {
								ck.Updates(map[string]interface{}{
									QQ:       0,
									Telegram: 0,
								})
								k++
							} else {
								b++
							}
						}
					}
					return fmt.Sprintf("操作成功，%d剩余许愿币%d", id, b)
				case "run", "执行":
					runTask(&Task{Path: v}, msgs...)
				}

			}
		}
		{
			o := false
			for _, v := range regexp.MustCompile(`京东账号\d*（(.*)）(.*)】(.*)`).FindAllStringSubmatch(msg, -1) {
				if !strings.Contains(v[3], "种子") && !strings.Contains(v[3], "undefined") {
					pt_pin := url.QueryEscape(v[1])
					for key, ss := range map[string][]string{
						"Fruit":        {"京东农场", "东东农场"},
						"Pet":          {"京东萌宠"},
						"Bean":         {"种豆得豆"},
						"JdFactory":    {"东东工厂"},
						"DreamFactory": {"京喜工厂"},
						"Jxnc":         {"京喜农场"},
						"Jdzz":         {"京东赚赚"},
						"Joy":          {"crazyJoy"},
						"Sgmh":         {"闪购盲盒"},
						"Cfd":          {"财富岛"},
						"Cash":         {"签到领现金"},
					} {
						for _, s := range ss {
							if strings.Contains(v[2], s) && v[3] != "" {
								if ck, err := GetJdCookie(pt_pin); err == nil {
									ck.Update(key, v[3])
								}
								if !o {
									o = true
								}
							}
						}
					}
				}
			}
			if o {
				return "导入互助码成功"
			}
		}
		for k, v := range replies {
			if regexp.MustCompile(k).FindString(msg) != "" {
				if regexp.MustCompile(`^https{0,1}://[^\x{4e00}-\x{9fa5}\n\r\s]{3,}$`).FindString(v) != "" {
					url := v
					rsp, err := httplib.Get(url).Response()
					if err != nil {
						return nil
					}
					return rsp
				}
				return v
			}
		}
	}
	return nil
}
