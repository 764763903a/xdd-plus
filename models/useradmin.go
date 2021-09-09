package models

import "github.com/beego/beego/v2/core/logs"

type UserAdmin struct {
	ID      int
	Content string `gorm:"unique"`
}

func IsUserAdmin(id string) bool {
	user := UserAdmin{}
	logs.Info(id)
	db.Model(UserAdmin{}).Where(Content+" = ?", id).First(user)
	logs.Info(user.Content)
	if len(user.Content) > 0 {
		return true
	}
	return false
}
