package service

import (
	. "example/homework/chatapp/utils"

	"github.com/spf13/viper"
)

type configName string

type MongoDBConfig struct {
	DBUsername    string  `mapstructure:"MONGODB_DB_USERNAME"`
	DBPassword    string  `mapstructure:"MONGODB_DB_PASSWORD"`
	DBHost        string  `mapstructure:"MONGODB_DB_HOST"`
	DBPort        string  `mapstructure:"MONGODB_DB_PORT"`
	DBMaxPoolSize *uint64 `mapstructure:"MONGODB_DB_OPTIONS_MAXPOOLSIZE"`
	DBName        string  `mapstructure:"MONGODB_DB_OPTIONS_DBNAME"`
	DBWritePolicy string  `mapstructure:"MONGODB_DB_OPTIONS_W"`
}

type LineConfig struct {
	LineSecret string `mapstructure:"LINEBOT_CH_SEC"`
	LineToken  string `mapstructure:"LINEBOT_CH_TOKEN"`
}

type AppConfig struct {
	Host string `mapstructure:"APP_HOST"`
	Port string `mapstructure:"APP_PORT"`
}

const (
	dbCfgName   configName = "mongodb"
	lineCfgName configName = "linebot"
	appCfgName  configName = "app"
)

func LoadDbCfg() *MongoDBConfig {
	cfg := &MongoDBConfig{}
	viper.AddConfigPath(getConfigPath())
	viper.SetConfigType("env")
	viper.SetConfigName(string(dbCfgName))
	if err := viper.ReadInConfig(); err != nil {
		Logger.Panicln("Failed to read config of db by viper", err)
	}
	viper.Unmarshal(&cfg)
	return cfg
}

func LoadLineCfg() *LineConfig {
	cfg := &LineConfig{}
	viper.AddConfigPath(getConfigPath())
	viper.SetConfigType("env")
	viper.SetConfigName(string(lineCfgName))
	if err := viper.ReadInConfig(); err != nil {
		Logger.Panicln("Failed to read config of line by viper", err)
	}
	viper.Unmarshal(&cfg)
	return cfg
}

func LoadAppCfg() *AppConfig {
	cfg := &AppConfig{}
	viper.AddConfigPath(getConfigPath())
	viper.SetConfigType("env")
	viper.SetConfigName(string(appCfgName))
	if err := viper.ReadInConfig(); err != nil {
		Logger.Panicln("Failed to read config of app by viper", err)
	}
	viper.Unmarshal(&cfg)
	return cfg
}
