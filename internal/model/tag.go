package model

import (
	"OceanQA/pkg/database"
	"OceanQA/pkg/logger"
)

type Tag struct {
	Id   uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;type:varchar(255);Index"`
}

func (t *Tag) TableName() string {
	return "tags"
}

func (t *Tag) Create() bool {
	if err := database.DB.Model(&Tag{}).Create(t).Error; err != nil {
		logger.WarnString("Database", "Create tag failed", err.Error())
		return false
	}
	return true
}
