package boot

import (
	"OceanQA/internal/conf"
	"OceanQA/pkg/config"
	"OceanQA/pkg/logger"
)

func Initialize() {
	config.InitConfig()
	conf.LoadConf()
	logger.InitLogger()
}
