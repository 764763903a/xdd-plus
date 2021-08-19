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
			// fmt.Println(m.Text, m.FromGroup())
			if !m.FromGroup() {
				rt := handleMessage(m.Text, "tg", m.Sender.ID)
				// fmt.Println(rt)
				switch rt.(type) {
				case string:
					b.Send(m.Sender, rt.(string))
				case *http.Response:
					b.SendAlbum(m.Sender, tb.Album{&tb.Photo{File: tb.FromReader(rt.(*http.Response).Body)}})
				}
			} else {
				rt := handleMessage(m.Text, "tgg", m.Sender.ID, int(m.Chat.ID), m.Sender)
				// fmt.Println(rt)
				switch rt.(type) {
				case string:
					b.Send(m.Chat, rt.(string), &tb.SendOptions{ReplyTo: m})
				case *http.Response:
					b.SendAlbum(m.Chat, tb.Album{&tb.Photo{File: tb.FromReader(rt.(*http.Response).Body)}}, &tb.SendOptions{ReplyTo: m})
				}
			}
		})
		logs.Info("监听tgbot")
		b.Start()
	}()
}

func SendTgMsg(uid int, msg string) {
	if b == nil || uid == 0 {
		return
	}
	b.Send(&tb.User{ID: uid}, msg)
}

func SendTggMsg(uid int, gid int, msg string) {
	if b == nil || uid == 0 {
		return
	}
	b.Send(&tb.Chat{ID: int64(gid)}, msg)
}
