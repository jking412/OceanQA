package boot

import (
	"OceanQA/internal/conf"
	"OceanQA/internal/model"
	"OceanQA/pkg/config"
	"OceanQA/pkg/database"
	"OceanQA/pkg/logger"
)

func Initialize() {
	config.InitConfig()
	conf.LoadConf()
	logger.InitLogger()
	database.InitDB()

	err := database.DB.AutoMigrate(&model.Question{},
		&model.Tag{})

	if err != nil {
		logger.FatalString("Database", "Database migration failed", err.Error())
		return
	}
}
