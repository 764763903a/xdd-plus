package models

import (
	"fmt"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Number   int `gorm:"unique"`
	Class    string
	ActiveAt time.Time
	Coin     int
}

func NewActiveUser(class string, uid int, msgs ...interface{}) {
	msg := ""
	if class == "tgg" {
		sender := msgs[4].(*tb.User)
		last := ""
		if sender.LastName != "" {
			last = " " + sender.LastName
		}
		if sender.Username == "" {
			msg = fmt.Sprintf(`@%s%s `, sender.FirstName, last)
		} else {
			msg = fmt.Sprintf(`@%s `, sender.Username)
		}

		class = "tg"
	}
	if class == "qqg" {
		class = "qq"
	}
	zero, _ := time.ParseInLocation("2006-01-02", time.Now().Local().Format("2006-01-02"), time.Local)
	var u User
	var ntime = time.Now()
	var first = false
	total := []int{}
	err := db.Where("class = ? and number = ?", class, uid).First(&u).Error
	if err != nil {
		first = true
		u = User{
			Class:    class,
			Number:   uid,
			Coin:     1,
			ActiveAt: ntime,
		}
		if err := db.Create(&u).Error; err != nil {
			return
		}
	} else {
		if zero.Unix() > u.ActiveAt.Unix() {
			first = true
			db.Updates(map[string]interface{}{
				"active_at": ntime,
				"coin":      gorm.Expr("coin+1"),
			})
			u.Coin++
		}
	}
	if first {
		db.Model(User{}).Select("count(id) as total").Where("active_at > ?", zero).Pluck("total", &total)
		msg += fmt.Sprintf("你是今天第%d个发言的用户，奖励%d个许愿币，许愿币余额%d。", total[0]+1, 1, u.Coin)
		// fmt.Println(msg)
		sendMessagee(msg, msgs...)
	}
}

func AddCoin(uid int) int {
	var u User
	db.Where("number = ?", uid).First(&u)
	db.Model(u).Updates(map[string]interface{}{
		"coin": gorm.Expr("coin+1"),
	})
	u.Coin++
	return u.Coin
}

func RemCoin(uid int) int {
	var u User
	db.Where("number = ?", uid).First(&u)
	db.Model(u).Updates(map[string]interface{}{
		"coin": gorm.Expr("coin-1"),
	})
	u.Coin--
	return u.Coin
}

func GetCoin(uid int) int {
	var u User
	db.Where("number = ?", uid).First(&u)
	return u.Coin
}
