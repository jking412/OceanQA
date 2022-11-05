package config

import (
	"fmt"
	viperlib "github.com/spf13/viper"
	"sync"
)

var once sync.Once

var internalViper *viperlib.Viper

func InitConfig() {
	once.Do(func() {
		internalViper = newViper()
	})
}

func newViper() *viperlib.Viper {
	viper := viperlib.New()
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败")
		panic(err)
	}
	viper.WatchConfig()
	viper.AutomaticEnv()
	return viper
}

func GetString(key string) string {
	return internalViper.GetString(key)
}

func GetInt(key string) int {
	return internalViper.GetInt(key)
}

func GetBool(key string) bool {
	return internalViper.GetBool(key)
}

func IsSet(key string) bool {
	return internalViper.IsSet(key)
}
