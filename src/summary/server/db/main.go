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

func (dbm *DBManger) ListSummaries(summary *Summary, offset int, limit int) (*[]Summary, error) {
	var summaries []Summary
	err := dbm.DB.Where(summary).Offset(offset).Limit(limit).Find(&summaries).Error
	return &summaries, err
}
