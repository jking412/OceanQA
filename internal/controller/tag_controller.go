package controller

import (
	"OceanQA/internal/model"
	"OceanQA/internal/request"
	"OceanQA/internal/response"
	"OceanQA/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	tagService = service.NewTagService()
)

type TagController struct {
}

func (*TagController) CreateTag(c *gin.Context) {
	req := &request.Tag{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "参数错误",
		})
		return
	}

	tag := &model.Tag{
		Name: req.Name,
	}

	if !tag.Create() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "创建标签失败",
		})
		return
	}

	tag, _ = tagService.GetTagByName(tag.Name)

	c.JSON(http.StatusOK, gin.H{
		"msg": &response.Tag{
			Id:   tag.Id,
			Name: tag.Name,
		},
	})

}

func (*TagController) SearchTag(c *gin.Context) {
	query := c.Query("query")
	tags, ok := tagService.QueryTags(query)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "查询标签失败",
		})
		return
	}
	resp := make([]response.Tag, 0)
	for _, tag := range tags {
		resp = append(resp, response.Tag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": resp,
	})
}
