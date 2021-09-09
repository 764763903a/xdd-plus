package models

type UserAdmin struct {
	ID      int
	Content string `gorm:"unique"`
}

func IsUserAdmin(id string) bool {
	user := &UserAdmin{}
	db.Where(Content+" = ?", id).First(user)
	if len(user.Content) > 0 {
		return true
	}
	return false
}

func RemoveUserAdmin(id string) bool {
	user := &UserAdmin{}
	db.Where(Content+" = ?", id).Delete(user)
	if len(user.Content) > 0 {
		return true
	}
	return false
}
