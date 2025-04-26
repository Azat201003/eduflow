package db

import (
	"fmt"

	"github.com/Azat201003/eduflow_service_api/config"
	"gorm.io/gorm"
)

type Summary struct {
	Title       string
	Description string
	FilePath    string
	ID          uint64
}

func (*Summary) TableName() string {
	conf, err := config.GetConfig("../../../../config.yaml")
	if err != nil {
		return "eduflow_summary.summaries"
	}
	serv, err := conf.GetServiceById(1)
	if err != nil {
		return "eduflow_summary.summaries"
	}
	return fmt.Sprintf("%v.summaries", serv.Connect.Schema)
}

func CreateSummary(db *gorm.DB, user *Summary) error {
	err := db.Create(user).Error
	return err
}

func FindSummary(db *gorm.DB, user *Summary) error {
	err := db.First(user, user).Error
	return err
}
