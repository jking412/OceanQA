package controller

import (
	"OceanQA/internal/response"
	"OceanQA/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	questionService = service.NewQuestionService()
)

type QuestionController struct {
}

func (*QuestionController) ShowAllQuestion(c *gin.Context) {
	currentNum, _ := strconv.ParseUint(c.Query("current_num"), 10, 64)
	pageSize, _ := strconv.ParseUint(c.Query("page_size"), 10, 64)
	var total uint64
	questions, ok := questionService.GetQuestionList()
	total = uint64(len(questions))
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "获取问题列表失败",
		})
		return
	}
	resp := &response.QuestionResponse{}
	if pageSize*currentNum <= total {
		questions = questions[pageSize*(currentNum-1) : pageSize*currentNum]
	} else if pageSize*(currentNum-1) <= total {
		questions = questions[pageSize*(currentNum-1):]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "页码超出范围",
		})
	}

	for index, question := range questions {
		resp.QuestionList = append(resp.QuestionList, response.Question{
			Id:        question.Id,
			CreatedAt: question.CreatedAt.Format("2006-01-02"),
			UpdatedAt: question.UpdatedAt.Format("2006-01-02"),
			Title:     question.Title,
			Content:   question.OriginContent,
			IsStared:  question.IsStared,
		})
		for _, tag := range question.Tags {
			resp.QuestionList[index].Tags = append(resp.QuestionList[index].Tags, tag.Name)
		}
	}
	resp.Pagination.PageNum = currentNum
	resp.Pagination.PageSize = pageSize
	resp.Pagination.Total = total
	c.JSON(http.StatusOK, gin.H{
		"msg": resp,
	})
}
