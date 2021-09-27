package models

import (
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type Limit struct {
	ID       int `gorm:"column:ID;primaryKey"`
	Number   int
	ActiveAt string
	Typ      int
	Num      int
}

func getLimit(uid int, typ int) bool {
	if Config.Lim == 0 {
		return true
	}
	u := &Limit{}
	err := db.Where("number = ? and typ = ? and active_at = ?", uid, typ, time.Now().Format("2006-01-02")).First(&u).Error
	if err == nil {
		logs.Info(u.Number)
		if u.Num < Config.Lim {
			db.Where("ID = ?", u.ID).Updates(&Limit{
				Num: u.Num + 1,
			})
			return true
		} else {
			return false
		}
	} else {
		begin := db.Begin()
		begin.Create(&Limit{
			ActiveAt: time.Now().Format("2006-01-02"),
			Typ:      typ,
			Number:   uid,
			Num:      1,
		})
		begin.Commit()
		return true
	}
}
