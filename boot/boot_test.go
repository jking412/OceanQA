package boot

import (
	"OceanQA/internal/model"
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestInit(t *testing.T) {
	Initialize()
	// 插入批量的Question数据
	for i := 0; i < 10; i++ {
		question := &model.Question{
			Title:         faker.Word(),
			AnswerUrl:     faker.URL(),
			OriginContent: faker.Paragraph(),
			ParseContent:  faker.Paragraph(),
			IsStared:      false,
		}
		question.Create()
	}
}
