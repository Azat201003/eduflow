package db

import (
	"math/rand"

	"gorm.io/gorm"
)

type User struct {
	Username string
	Password string
	Token    string
	ID       uint64
}

func (*User) TableName() string {
	return "users.users"
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateToken() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CreateUser(db *gorm.DB, user *User) error {
	user.Token = generateToken()
	err := db.Create(user).Error
	return err
}

func FindUser(db *gorm.DB, user *User) error {
	err := db.Find(user).Error
	return err
}
