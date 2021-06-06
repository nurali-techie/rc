package config

import "github.com/spf13/viper"

const (
	DB_PATH = "DB_PATH"
)

type config struct {
	dbPath string
}

var cfg *config

func Init() {
	viperInit()
	cfg = &config{}

	cfg.dbPath = viper.GetString(DB_PATH)
}

func viperInit() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func GetDBPath() string {
	return cfg.dbPath
}
