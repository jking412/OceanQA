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
