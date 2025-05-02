package db

import (
	"errors"
	"fmt"

	"github.com/Azat201003/eduflow_service_api/config"
	"gorm.io/gorm"
)

type Summary struct {
	Title       string
	Description string
	FilePath    string
	ID          uint64
	AuthorId    uint64
}

type DBManger struct {
	DB *gorm.DB
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

func (dbm *DBManger) CreateSummary(summary *Summary) error {
	err := dbm.DB.Create(summary).Error
	return err
}

func (dbm *DBManger) FindSummary(summary *Summary) error {
	r := dbm.DB.First(summary, summary)
	if r.RowsAffected == 0 {
		return errors.New("Find no records")
	}
	err := r.Error
	return err
}

func UpdateSummary(db *gorm.DB, summary *Summary) error {
	err := db.Save(summary).Error
	return err
}

func DeleteSummary(db *gorm.DB, summary *Summary) error {
	err := db.Delete(summary).Error
	return err
}

func ListSummaries(db *gorm.DB, summaries *[]Summary) error {
	err := db.Find(summaries).Error
	return err
}
