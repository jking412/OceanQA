package database

import (
	"OceanQA/internal/conf"
	"OceanQA/pkg/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once

var DB *gorm.DB

func InitDB() {
	once.Do(func() {
		DB = ConnectDB()
	})
}

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBConf.Username,
		conf.DBConf.Password,
		conf.DBConf.Host,
		conf.DBConf.Port,
		conf.DBConf.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.FatalString("Database", "Database connection failed. DSN"+
			": "+dsn, err.Error())
		return nil
	}

	logger.InfoString("Database", "Database connected", conf.DBConf.DBName)

	return db
}
