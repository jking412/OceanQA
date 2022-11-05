package service

import (
	"OceanQA/internal/model"
	"OceanQA/pkg/database"
)

type TagService struct {
}

func NewTagService() *TagService {
	return &TagService{}
}

func (ts *TagService) QueryTags(query string) ([]model.Tag, bool) {
	tags := make([]model.Tag, 0)
	if err := database.DB.Model(&model.Tag{}).Where("name LIKE ?", "%"+query+"%").Find(&tags).Error; err != nil {
		return tags, false
	}
	return tags, true
}

func (ts *TagService) GetTagByName(name string) (*model.Tag, bool) {
	tag := &model.Tag{}
	if err := database.DB.Model(&model.Tag{}).Where("name = ?", name).First(tag).Error; err != nil {
		return tag, false
	}
	return tag, true
}
