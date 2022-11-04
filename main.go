package main

import (
	"OceanQA/boot"
	"OceanQA/internal/conf"
	"OceanQA/pkg/template"
	"OceanQA/router"
	"github.com/gin-gonic/gin"
)

func main() {
	boot.Initialize()

	r := gin.Default()

	router.RegisterRoutes(r)

	template.Load(r)

	r.Run(":" + conf.ServerConf.Port)
}
