package db

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Azat201003/eduflow/backend/libs/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

type NoParamsError struct{}

func (err *NoParamsError) GRPCStatus() *status.Status {
	return status.New(codes.Canceled, err.Error())
}
func (*NoParamsError) Error() string { return "No params given." }

type NotFoundError struct{}

func (err *NotFoundError) GRPCStatus() *status.Status {
	return status.New(codes.Canceled, err.Error())
}
func (*NotFoundError) Error() string { return "No params given." }

func (dbm *DBManger) FindUser(user *User) error {
	if *user == (User{}) {
		return errors.New("No params")
	}
	r := dbm.DB.First(user, user)
	if r.RowsAffected == 0 {
		return errors.New("Found no records")
	}
	return r.Error
}
