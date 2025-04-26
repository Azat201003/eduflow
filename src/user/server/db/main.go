package db

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Azat201003/eduflow_service_api/config"
	"gorm.io/gorm"
)

type User struct {
	Username string
	Password string
	Token    string
	ID       uint64
}

func (*User) TableName() string {
	conf, err := config.GetConfig("../../../../config.yaml")
	if err != nil {
		return "users.users"
	}
	serv, err := conf.GetServiceById(0)
	if err != nil {
		return "users.users"
	}
	return fmt.Sprintf("%v.users", serv.Connect.Schema)
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
	//  RowsAffected
	r := db.First(user, user)
	if r.RowsAffected == 0 {
		return errors.New("Find no records")
	}
	err := r.First(user).Error
	return err
}
