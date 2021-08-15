package models

func (ck *JdCookie) Push(msg string) {
	if Config.QywxKey != "" {
		go qywxNotify(&QywxConfig{Content: msg})
	}
	if Config.TelegramBotToken != "" {
		go tgBotNotify(msg)
	}
	if Config.QQID != 0 {
		go SendQQ(Config.QQID, msg)
		go SendQQ(int64(ck.QQ), msg)
	}
	if ck.PushPlus != "" {
		go pushPlus(ck.PushPlus, msg)
	}
}
