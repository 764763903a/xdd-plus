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
var SendQQGroup = func(a int64, b int64, c interface{}) {

}
var ListenQQPrivateMessage = func(uid int64, msg string) {
	SendQQ(uid, handleMessage(msg, "qq", int(uid)))
}

var ListenQQGroupMessage = func(gid int64, uid int64, msg string) {
	if gid == Config.QQGroupID {
		if Config.QbotPublicMode {
			SendQQGroup(gid, uid, handleMessage(msg, "qqg", int(uid), int(gid)))
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
	uid := msgs[2].(int)
	gid := 0
	if len(msgs) >= 4 {
		gid = msgs[3].(int)
	}
	switch tp {
	case "tg":
		SendTgMsg(uid, msg)
	case "tgg":
		SendTggMsg(gid, uid, msg)
	case "qq":
		SendQQ(int64(uid), msg)
	case "qqg":
		SendQQGroup(int64(gid), int64(uid), msg)
	}
}

var isAdmin = func(msgs ...interface{}) bool {
	if len(msgs) == 0 {
		return false
	}
	tp := msgs[1].(string)
	uid := msgs[2].(int)
	switch tp {
	case "tg", "tgg":
		if int(Config.TelegramUserID) == uid {
			return true
		}
	case "qq", "qqg":
		if int(Config.QQID) == uid {
			return true
		}
	}
	return false
}

var handleMessage = func(msgs ...interface{}) interface{} {
	msg := msgs[0].(string)
	tp := msgs[1].(string)
	uid := msgs[2].(int)
	gid := 0
	if len(msgs) >= 4 {
		gid = msgs[3].(int)
	}

	switch msg {
	case "取消屏蔽":
		if !isAdmin(msgs...) {
			return "你没有权限操作"
		}
		e := db.Model(JdCookie{}).Where(fmt.Sprintf("%s != ?", Hack), False).Update(Hack, False).RowsAffected
		Save <- &JdCookie{}
		return fmt.Sprintf("操作成功，更新%d条记录", e)
	case "status", "状态":
		if !isAdmin(msgs...) {
			return "你没有权限操作"
		}
		return Count()
	case "打卡", "签到", "sign":
		NewActiveUser(tp, uid, msgs...)
	case "许愿币":
		return fmt.Sprintf("余额%d", GetCoin(uid))
	case "qrcode", "扫码", "二维码", "scan":
		url := fmt.Sprintf("http://127.0.0.1:%d/api/login/qrcode.png?tp=%s&uid=%d&gid=%d", web.BConfig.Listen.HTTPPort, tp, uid, gid)
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
	case "重启", "reload", "restart", "reboot":
		if !isAdmin(msgs...) {
			return "你没有权限操作"
		}
		sendMessagee("小滴滴重启程序", msgs...)
		Daemon()
		return nil
	case "任务列表":
		rt := ""
		for i := range Config.Repos {
			for j := range Config.Repos[i].Task {
				rt += fmt.Sprintf("%s\t%s\n", Config.Repos[i].Task[j].Title, Config.Repos[i].Task[j].Cron)
			}
		}
		return rt
	case "查询", "query":
		cks := GetJdCookies()
		tmp := []JdCookie{}
		for _, ck := range cks {
			if tp == "qq" || tp == "qqg" {
				if ck.QQ == uid {
					tmp = append(tmp, ck)
				}
			} else if tp == "tg" || tp == "tgg" {
				if ck.Telegram == uid {
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
			ss := regexp.MustCompile(`packetId=(\S+)(&|&amp;)currentActId`).FindStringSubmatch(msg)
			if len(ss) > 0 {
				if Cdle {
					return "推毛线啊"
				}
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
						if tp == "qq" || tp == "qqg" {
							ck.QQ = uid
						} else if tp == "tg" || tp == "tgg" {
							ck.Telegram = uid
						}
						if HasKey(ck.PtKey) {
							sendMessagee(fmt.Sprintf("作弊，许愿币-1，余额%d", RemCoin(uid)), msgs...)
						} else {
							if nck, err := GetJdCookie(ck.PtPin); err == nil {
								nck.InPool(ck.PtKey)
								msg := fmt.Sprintf("更新账号，%s", ck.PtPin)
								(&JdCookie{}).Push(msg)
								logs.Info(msg)
							} else {
								if Cdle {
									ck.Hack = True
								}
								NewJdCookie(&ck)
								msg := fmt.Sprintf("添加账号，%s", ck.PtPin)
								sendMessagee(fmt.Sprintf("很棒，许愿币+1，余额%d", AddCoin(uid)), msgs...)
								logs.Info(msg)
							}
						}
					} else {
						sendMessagee(fmt.Sprintf("无效，许愿币-1，余额%d", RemCoin(uid)), msgs...)
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
					b := 0
					for _, ck := range GetJdCookies() {
						if uid == ck.QQ || uid == ck.Telegram {
							b++
						}
					}
					if b < 5 {
						return "许愿币不足，需要5个许愿币。"
					} else {
						(&JdCookie{}).Push(fmt.Sprintf("%d许愿%s，许愿币余额%d。", uid, v, b))
						return "收到许愿，愿望达成后会自动扣除5个许愿币。"
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
					if !isAdmin(msgs...) {
						return "你没有权限操作"
					}
					runTask(&Task{Path: v}, msgs...)
				case "cmd", "command":
					if !isAdmin(msgs...) {
						return "你没有权限操作"
					}
					cmd(v, msgs...)
				}

			}
		}
		{
			o := false
			for _, v := range regexp.MustCompile(`京东账号\d*（(.*)）(.*)】(\S*)`).FindAllStringSubmatch(msg, -1) {
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
					ctp := rsp.Header.Get("content-type")
					if ctp == "" {
						rsp.Header.Get("Content-Type")
					}
					if strings.Contains(ctp, "text") || strings.Contains(ctp, "json") {
						data, _ := ioutil.ReadAll(rsp.Body)
						return string(data)
					}
					return rsp
				}
				return v
			}
		}
	}
	return nil
}
