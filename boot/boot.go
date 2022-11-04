package boot

import (
	"OceanQA/internal/conf"
	"OceanQA/pkg/config"
)

func Initialize() {
	config.InitConfig()
	conf.LoadConf()
}
