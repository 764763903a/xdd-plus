package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int
	Number   int
	Class    string
	ActiveAt time.Time
	Coin     int
}

func NewActiveUser(class string, uid int) (bool, string) {
	if class == "tgg" {
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
		if err := db.Create(u).Error; err != nil {
			return true, err.Error()
		}
	} else {
		if zero.After(u.ActiveAt) {
			first = true
			db.Updates(map[string]interface{}{
				"active_at": ntime,
				"coin":      gorm.Expr("coin+1"),
			})
			u.Coin++
		}
	}
	if first {
		db.Select("count(id) as total").Where("active_at > ?", u.ActiveAt).Pluck("total", &total)
	}
	return first, fmt.Sprintf("你是今天第%d个发言的用户，奖励%d个心愿币，心愿币余额%d", total[0]+1, 1, u.Coin)
}
