package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int
	Number   int `gorm:"unique"`
	Class    string
	ActiveAt time.Time
	Coin     int
}

func AddCoin(uid int) int {
	var u User
	if db.Where("number = ?", uid).First(&u).Error != nil {
		return 0
	}
	db.Model(u).Updates(map[string]interface{}{
		"coin": gorm.Expr("coin+1"),
	})
	u.Coin++
	return u.Coin
}

func RemCoin(uid int, num int) int {
	var u User
	db.Where("number = ?", uid).First(&u)
	db.Model(u).Updates(map[string]interface{}{
		"coin": gorm.Expr(fmt.Sprintf("coin-%d", num)),
	})
	u.Coin -= num
	return u.Coin
}

func GetCoin(uid int) int {
	var u User
	db.Where("number = ?", uid).First(&u)
	return u.Coin
}
