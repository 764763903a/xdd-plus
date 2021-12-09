package models

import (
	"github.com/Mrs4s/MiraiGo/utils"
	"io"
	"io/ioutil"
	"os"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Containers []Container
	// Tasks              []Task
	Qrcode              string
	Master              string
	Mode                string
	Static              string
	Database            string
	QywxKey             string `yaml:"qywx_key"`
	Resident            string
	UserAgent           string `yaml:"user_agent"`
	Theme               string
	TelegramBotToken    string `yaml:"telegram_bot_token"`
	TelegramUserID      int    `yaml:"telegram_user_id"`
	QQID                int64  `yaml:"qquid"`
	QQGroupID           int64  `yaml:"qqgid"`
	DefaultPriority     int    `yaml:"default_priority"`
	NoGhproxy           bool   `yaml:"no_ghproxy"`
	QbotPublicMode      bool   `yaml:"qbot_public_mode"`
	DailyAssetPushCron  string `yaml:"daily_asset_push_cron"`
	Version             string `yaml:"version"`
	CTime               string `yaml:"AtTime"`
	IsHelp              bool   `yaml:"IsHelp"`
	IsOldV4             bool   `yaml:"IsOldV4"`
	ApiToken            string `yaml:"ApiToken"`
	Wskey               bool   `yaml:"Wskey"`
	TGURL               string `yaml:"TGURL"`
	SMSAddress          string `yaml:"SMSAddress"`
	IsAddFriend         bool   `yaml:"IsAddFriend"`
	Lim                 int    `yaml:"Lim"`
	Tyt                 int    `yaml:"Tyt"`
	IFC                 bool   `yaml:"IFC"`
	Later               int    `yaml:"Later"`
	Node                string
	Npm                 string
	Python              string
	Pip                 string
	NoAdmin             bool   `yaml:"no_admin"`
	QbotConfigFile      string `yaml:"qbot_config_file"`
	Repos               []Repo
	HttpProxyServerPort int `yaml:"http_proxy_server_port"`
}

var Balance = "balance"
var Parallel = "parallel"
var GhProxy = "https://ghproxy.com/"
var Cdle = false

var Config Yaml

func initConfig() {
	confDir := ExecPath + "/conf"
	if _, err := os.Stat(confDir); err != nil {
		os.MkdirAll(confDir, os.ModePerm)
	}
	for _, name := range []string{"app.conf", "config.yaml", "reply.php", "qbot.yaml"} {
		f, err := os.OpenFile(ExecPath+"/conf/"+name, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			logs.Warn(err)
		}
		s, _ := ioutil.ReadAll(f)
		if len(s) == 0 {
			logs.Info("下载配置%s", name)
			r, err := httplib.Get(GhProxy + "https://raw.githubusercontent.com/764763903a/xdd-plus/main/conf/demo_" + name).Response()
			if err == nil {
				io.Copy(f, r.Body)
			}
		}
		f.Close()
	}
	content, err := ioutil.ReadFile(ExecPath + "/conf/config.yaml")
	if err != nil {
		logs.Warn("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &Config) != nil {
		logs.Warn("解析config.yaml出错: %v", err)
	}
	if ExecPath == "/Users/cdle/Desktop/xdd" || Config.NoAdmin {
		Cdle = true
	}
	if Config.Master == "" {
		Config.Master = "xxxx"
	}
	if Config.CTime == "" {
		Config.CTime = "10"
	}
	if Config.Mode != Parallel {
		Config.Mode = Balance
	}
	if Config.Qrcode != "" {
		Config.Theme = Config.Qrcode
	}
	if Config.NoGhproxy {
		GhProxy = ""
	}
	if Config.Tyt == 0 {
		Config.Tyt = 8
	}
	if Config.Later == 0 {
		Config.Later = 60
	}
	if Config.Database == "" {
		Config.Database = ExecPath + "/.xdd.db"
	}
	if Config.Npm == "" {
		Config.Npm = "npm"
	}
	if Config.ApiToken == "" {
		Config.ApiToken = utils.RandomString(17)
	}
	if Config.Node == "" {
		Config.Node = "node"
	}
	if Config.Python == "" {
		Config.Python = "python3"
	}
	if Config.Pip == "" {
		Config.Pip = "pip3"
	}
}
