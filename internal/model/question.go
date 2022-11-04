package model

import (
	"OceanQA/pkg/database"
	"OceanQA/pkg/logger"
	"time"
)

type Question struct {
	Id            uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime"`
	Title         string    `gorm:"column:title;type:varchar(255);Index"`
	AnswerUrl     string    `gorm:"column:answer_url;type:varchar(255)"`
	OriginContent string    `gorm:"column:content;type:longtext"`
	ParseContent  string    `gorm:"column:parse_content;type:longtext"`
	IsStared      bool      `gorm:"column:is_stared;type:tinyint(1)"`
	Tag           []Tag     `gorm:"many2many:question_tags;"`
}

func (q *Question) TableName() string {
	return "questions"
}

func (q *Question) Create() bool {
	if err := database.DB.Model(&Question{}).Create(q).Error; err != nil {
		logger.WarnString("Database", "Create question failed", err.Error())
		return false
	}
	return true
}
