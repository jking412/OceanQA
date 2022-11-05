package boot

import (
	"OceanQA/internal/model"
	"OceanQA/internal/service"
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestCreateQuestion(t *testing.T) {
	Initialize()
	// 插入批量的Question数据
	for i := 0; i < 1; i++ {
		question := &model.Question{
			Title:         faker.Word(),
			AnswerUrl:     faker.URL(),
			OriginContent: faker.Paragraph(),
			ParseContent:  faker.Paragraph(),
			IsStared:      false,
			Tags: []model.Tag{
				{
					Id: 1,
				},
				{
					Id: 2,
				},
			},
		}
		question.Create()
	}
}

func TestCreateTag(t *testing.T) {
	Initialize()
	// 插入批量的Tag数据
	for i := 0; i < 10; i++ {
		tag := &model.Tag{
			Name: faker.Word(),
		}
		tag.Create()
	}
}

func TestCreateQuestionTag(t *testing.T) {
	Initialize()
	questionService := &service.QuestionService{}
	// 插入批量的QuestionTag数据
	for i := 0; i < 10; i++ {
		m := &model.QuestionTag{
			QuestionId: uint64(i + 1),
			TagId:      uint64(i + 1),
		}
		questionService.CreateQuestionTag(m.QuestionId, m.TagId)
	}
}
