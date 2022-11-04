package main

import (
	"OceanQA/boot"
	"OceanQA/internal/conf"
	"OceanQA/pkg/logger"
	"OceanQA/pkg/template"
	"OceanQA/router"
	"github.com/gin-gonic/gin"
)

func main() {
	boot.Initialize()

	r := gin.Default()

	router.RegisterRoutes(r)

	template.Load(r)

	logger.Info("Server is running at " + ":" + conf.ServerConf.Port)

	if err := r.Run(":" + conf.ServerConf.Port); err != nil {
		logger.FatalString("Server", "Server start failed", err.Error())
		return
	}
}
