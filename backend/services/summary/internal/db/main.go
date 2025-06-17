package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/Azat201003/eduflow/backend/libs/config"
	pq "github.com/lib/pq"
	"gorm.io/gorm"
)

type Summary struct {
	Title       string
	Description string
	FilePath    string
	ID          uint64
	AuthorId    uint64
	Tags        pq.Int64Array `gorm:"type:integer[]"`
}

type Tag struct {
	ID    uint64
	Name  string
	Color string
}

type DBManger struct {
	DB *gorm.DB
}

func (*Summary) TableName() string {
	conf, err := config.GetConfig("../../../config.yaml")
	if err != nil {
		return "eduflow_summary.summaries"
	}
	serv, err := conf.GetServiceById(1)
	if err != nil {
		return "eduflow_summary.summaries"
	}
	return fmt.Sprintf("%v.tags", serv.Connect.Schema)
}

func (*Tag) TableName() string {
	conf, err := config.GetConfig("../../../config.yaml")
	if err != nil {
		return "eduflow_summary.tags"
	}
	serv, err := conf.GetServiceById(1)
	if err != nil {
		return "eduflow_summary.tags"
	}
	return fmt.Sprintf("%v.tags", serv.Connect.Schema)
}

// summaries

func (dbm *DBManger) CreateSummary(summary *Summary) error {
	log.Println(summary)
	err := dbm.DB.Create(summary).Error
	return err
}

func (dbm *DBManger) FindSummary(summary *Summary) error {
	fmt.Println(len(summary.Tags))
	tags := summary.Tags
	summary.Tags = nil
	fmt.Println(len(tags))
	r := dbm.DB.First(summary, summary)
	if r.RowsAffected == 0 {
		return errors.New("Find no records")
	}
	err := r.Error
	return err
}

func (dbm *DBManger) FilteredSummaries(summary *Summary, offset int, limit int) (*[]Summary, error) {
	var summaries []Summary
	fmt.Println(summary.Tags)
	tags := summary.Tags
	summary.Tags = nil
	fmt.Println(tags)
	err := dbm.DB.Where(summary).Where("tags @> ?", tags).Offset(offset).Limit(limit).Find(&summaries).Error
	fmt.Println(summaries)
	return &summaries, err
}

// tags

func (dbm *DBManger) FindTag(tag *Tag) error {
	r := dbm.DB.First(tag, tag)
	if r.RowsAffected == 0 {
		return errors.New("Find no records")
	}
	err := r.Error
	return err
}

func (dbm *DBManger) FilteredTags(filter *Tag, offset int, limit int) (*[]Tag, error) {
	var tags []Tag
	err := dbm.DB.Where(filter).Offset(offset).Limit(limit).Find(&tags).Error
	return &tags, err
}
