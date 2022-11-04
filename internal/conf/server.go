package conf

import "OceanQA/pkg/config"

var ServerConf = struct {
	Port string
}{}

func loadServerConf() {
	ServerConf.Port = config.LoadString("server.port", "9999")
}
