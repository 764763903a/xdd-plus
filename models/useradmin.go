package models

type UserAdmin struct {
	ID      int
	Content string `gorm:"unique"`
}

func IsUserAdmin(id string) bool {
	user := UserAdmin{}
	db.Model(UserAdmin{}).Where(Content+" = ?", id).First(user)
	if len(user.Content) > 0 {
		return true
	}
	return false
}
