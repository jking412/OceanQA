package template

import (
	"OceanQA/ui"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(r *gin.Engine) {
	r.StaticFS("/static", http.FS(&ui.StaticResource{
		FS: ui.Dist,
	}))

	r.LoadHTMLGlob("ui/dist/*.html")
}
