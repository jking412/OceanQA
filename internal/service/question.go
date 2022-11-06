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

func (qs *QuestionService) GetQuestion(id uint64) (*model.Question, bool) {
	question := &model.Question{}
	if err := database.DB.Model(&model.Question{}).Preload("Tags").First(question, id).Error; err != nil {
		logger.WarnString("Database", "Get question failed", err.Error())
		return question, false
	}
	question.ReadTime++
	if !qs.UpdateQuestionReadTime(question.Id, question.ReadTime) {
		logger.WarnString("Database", "Get question failed", "Update question read time failed")
		return question, false
	}
	return question, true
}

func (qs *QuestionService) UpdateQuestionReadTime(id uint64, readTime uint64) bool {
	if err := database.DB.Model(&model.Question{}).Where("id = ?", id).Update("read_time", readTime).Error; err != nil {
		logger.WarnString("Database", "Update question read time failed", err.Error())
		return false
	}
	return true
}

func (qs *QuestionService) UpdateQuestionStarStatus(id uint64, isStared bool) bool {
	if err := database.DB.Model(&model.Question{}).Where("id = ?", id).Update("is_stared", isStared).Error; err != nil {
		logger.WarnString("Database", "Update question star status failed", err.Error())
		return false
	}
	return true
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

func (qs *QuestionService) DeleteQuestion(id uint64) bool {
	if err := database.DB.Table("question_tags").Delete(&model.QuestionTag{}, "question_id = ?", id).Error; err != nil {
		logger.WarnString("Database", "Delete question tag failed", err.Error())
		return false
	}
	if err := database.DB.Delete(&model.Question{Id: id}).Error; err != nil {
		logger.WarnString("Database", "Delete question failed", err.Error())
		return false
	}
	return true
}
