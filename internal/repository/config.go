package repository

import "github.com/spf13/viper"

type Config struct {
	HOST   string `yaml:"host"`
	USER   string `yaml:"user"`
	DBNAME string `yaml:"dbname"`
	PORT   string `yaml:"port"`
	PASS   string `yaml:"pass"`
}

func getConfig() Config {
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	if viper.ReadInConfig() != nil {
		panic(viper.ReadInConfig())
	}
	return Config{
		HOST:   viper.GetString("db.host"),
		PORT:   viper.GetString("db.port"),
		DBNAME: viper.GetString("db.dbname"),
		USER:   viper.GetString("db.user"),
		PASS:   viper.GetString("db.pass"),
	}
}
