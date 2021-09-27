package models

import "time"

type Limit struct {
	ID       int `gorm:"column:ID;primaryKey"`
	Number   int
	ActiveAt string
	Typ      int
	Num      int
}

func getLimit(uid int, typ int) bool {
	u := &Limit{}
	db.Where("number = ? and typ = ? and ActiveAt = ?", uid, typ, time.Now().Format("2006-01-02")).First(&u)
	if Config.Lim == 0 {
		return true
	}
	if u.ID != 0 {
		if u.Num < Config.Lim {
			db.Updates(&Limit{
				Num: u.Num + 1,
			}).Where("ID = ?", u.ID)
			return true
		} else {
			return false
		}
	} else {
		u.ActiveAt = time.Now().Format("2006-01-02")
		u.Typ = typ
		u.Number = uid
		u.Num = 1
		db.Create(u)
		return true
	}
}
