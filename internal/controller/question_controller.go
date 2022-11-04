package controller

import (
	"OceanQA/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	questionService = service.NewQuestionService()
)

type QuestionController struct {
}

func (*QuestionController) ShowAllQuestion(c *gin.Context) {
	questions, ok := questionService.GetQuestionList()
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "获取问题列表失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": questions,
	})
}
