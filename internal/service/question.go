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
	if err := database.DB.Model(&model.Question{}).Find(&questions).Error; err != nil {
		logger.WarnString("Database", "GetQuestionList Failed", err.Error())
		return questions, false
	}
	return questions, true
}
