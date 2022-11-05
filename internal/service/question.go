package service

import (
	"OceanQA/internal/model"
	"OceanQA/pkg/database"
	"OceanQA/pkg/logger"
)

type QuestionService struct {
}

func NewQuestionService() *QuestionService {
	return &QuestionService{}
}

func (qs *QuestionService) GetQuestionList() ([]model.Question, bool) {
	questions := make([]model.Question, 0)
	if err := database.DB.Model(&model.Question{}).Preload("Tags").Find(&questions).Error; err != nil {
		logger.WarnString("Database", "GetQuestionList Failed", err.Error())
		return questions, false
	}
	return questions, true
}

func (qs *QuestionService) CreateQuestionTag(questionId uint64, tagId uint64) bool {
	questionTag := model.QuestionTag{
		QuestionId: questionId,
		TagId:      tagId,
	}
	if err := database.DB.Model(&model.QuestionTag{}).Create(&questionTag).Error; err != nil {
		logger.WarnString("Database", "Create question tag failed", err.Error())
		return false
	}
	return true
}
