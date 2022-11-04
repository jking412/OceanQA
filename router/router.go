package router

import (
	"OceanQA/ui"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
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
