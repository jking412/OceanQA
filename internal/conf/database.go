package conf

import "OceanQA/pkg/config"

var DBConf = struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}{}

func loadDBConf() {
	DBConf.Host = config.LoadString("database.host", "localhost")
	DBConf.Port = config.LoadInt("database.port", 3306)
	DBConf.Username = config.LoadString("database.username", "root")
	DBConf.Password = config.LoadString("database.password", "123456")
	DBConf.DBName = config.LoadString("database.dbname", "ocean_qa")
}
