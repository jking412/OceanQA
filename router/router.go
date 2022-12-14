package router

import (
	"OceanQA/internal/controller"
	"OceanQA/internal/middleware"
	"OceanQA/ui"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	questionController = &controller.QuestionController{}
	tagController      = &controller.TagController{}
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(middleware.Cors())

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	v1Group := r.Group("/v1")
	{
		questionGroup := v1Group.Group("/question")
		{
			questionGroup.GET("/all", questionController.ShowAllQuestion)
			questionGroup.POST("/create", questionController.CreateQuestion)
			questionGroup.GET("/show", questionController.ShowQuestion)
			questionGroup.GET("/delete", questionController.DeleteQuestion)
			questionGroup.GET("/star", questionController.UpdateStarStatus)
		}

		tagGroup := v1Group.Group("/tag")
		{
			tagGroup.POST("/create", tagController.CreateTag)
			tagGroup.GET("/search", tagController.SearchTag)
		}
	}

	r.Static("/css", "ui/dist/css")
	r.Static("/js", "ui/dist/js")
	r.Static("/fonts", "ui/dist/fonts")

	r.NoRoute(NoRouteHandler)
}

func NoRouteHandler(c *gin.Context) {
	name := c.Request.URL.Path
	if name != "/favicon.ico" && name != "/favicon.ico/" {
		c.Status(http.StatusNotFound)
		return
	}
	filePath := "dist" + name
	var file []byte
	var err error
	file, err = ui.Dist.ReadFile(filePath)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.String(http.StatusOK, string(file))
}
