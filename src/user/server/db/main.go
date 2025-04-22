package db

import "gorm.io/gorm"

type User struct {
	Username string
	Password string
	Token    string
	ID       uint64
}

func generateToken() string {
	return "abeme"
}

func CreateUser(db *gorm.DB, user *User) {
	err := db.Create(User{
		Username: user.Username,
		Password: user.Password,
		Token: generateToken()
	}).Error
}
