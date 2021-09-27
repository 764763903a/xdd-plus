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
	if Config.Lim == 0 {
		return true
	}
	u := &Limit{}
	err := db.Where("number = ? and typ = ? and active_at = ?", uid, typ, time.Now().Format("2006-01-02")).First(&u).Error
	if err != nil {
		if u.Num < Config.Lim {
			db.Updates(&Limit{
				Num: u.Num + 1,
			}).Where("ID = ?", u.ID)
			return true
		} else {
			return false
		}
	} else {
		db.Begin()
		u.ActiveAt = time.Now().Format("2006-01-02")
		u.Typ = typ
		u.Number = uid
		u.Num = 1
		db.Create(u)
		db.Commit()
		return true
	}
}
