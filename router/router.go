package router

import (
	"OceanQA/internal/controller"
	"OceanQA/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	questionController = &controller.QuestionController{}
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(middleware.Cors())

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

	r.Static("/css", "ui/dist/css")
	r.Static("/js", "ui/dist/js")
	r.Static("/fonts", "ui/dist/fonts")
	r.Static("/favicon.ico", "ui/dist")

	r.NoRoute(NoRouteHandler)
}

func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"msg": "error",
	})

}
