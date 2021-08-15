package models

import (
	"net/http"
	"time"

	"github.com/beego/beego/v2/core/logs"
	tb "gopkg.in/tucnak/telebot.v2"
)

var b *tb.Bot

func initTgBot() {
	go func() {
		if Config.TelegramBotToken == "" {
			return
		}
		var err error
		b, err = tb.NewBot(tb.Settings{
			Token:  Config.TelegramBotToken,
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			logs.Warn("监听tgbot失败")
			return
		}
		b.Handle(tb.OnText, func(m *tb.Message) {
			rt := handleMessage(m.Text, "tg", m.Sender.ID)
			switch rt.(type) {
			case string:
				b.Send(m.Sender, rt.(string))
			case *http.Response:
				b.SendAlbum(m.Sender, tb.Album{&tb.Photo{File: tb.FromReader(rt.(*http.Response).Body)}})
			}
		})
		logs.Info("监听tgbot")
		b.Start()
	}()
}

func tgBotNotify(msg string) {
	if b == nil {
		return
	}
	if Config.TelegramUserID == 0 {
		logs.Warn("tgbot未绑定用id")
		return
	}
	b.Send(&tb.User{ID: Config.TelegramUserID}, msg)
}

func SendTgMsg(id int, msg string) {
	if b == nil || id == 0 {
		return
	}
	b.Send(&tb.User{ID: id}, msg)
}
