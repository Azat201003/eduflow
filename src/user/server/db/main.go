package db

import (
	"errors"
	"fmt"
	"log"
	"math/rand"

	"github.com/Azat201003/eduflow_service_api/config"
	"gorm.io/gorm"
)

type User struct {
	Username string
	Password string
	Token    string
	ID       uint64
	IsStaff  bool
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

type DBManger struct {
	DB *gorm.DB
}

func (dbm *DBManger) CreateUser(user *User) error {
	user.Token = generateToken()
	err := dbm.DB.Create(user).Error
	return err
}

func (dbm *DBManger) FindUser(user *User) error {
	log.Println(user.ID)
	r := dbm.DB.First(user, user)
	log.Println(r.Error, r.RowsAffected)
	if r.RowsAffected == 0 {
		return errors.New("Find no records")
	}
	return r.Error
}
