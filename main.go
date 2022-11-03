package main

import (
	"OceanQA/ui"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.StaticFS("/static", http.FS(&ui.StaticResource{
		FS: ui.Dist,
	}))

	r.LoadHTMLGlob("ui/dist/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.NoRoute(func(c *gin.Context) {
		requestUri := c.Request.RequestURI
		file, err := ui.Dist.ReadFile("dist" + requestUri)
		if err != nil {
			c.Abort()
		} else {
			c.String(http.StatusOK, string(file))
		}
	})

	r.Run(":8080")
}
