package router

import (
	"OceanQA/internal/controller"
	"OceanQA/ui"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	questionController = &controller.QuestionController{}
)

func RegisterRoutes(r *gin.Engine) {

	v1Group := r.Group("/v1")
	{
		questionGroup := v1Group.Group("/question")
		{
			questionGroup.GET("/all", questionController.ShowAllQuestion)
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.NoRoute(staticHandler)
}

func staticHandler(c *gin.Context) {
	requestUri := c.Request.RequestURI
	file, err := ui.Dist.ReadFile("dist" + requestUri)
	if err != nil {
		c.Abort()
	} else {
		c.String(http.StatusOK, string(file))
	}
}
