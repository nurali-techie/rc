package cfg

import "github.com/spf13/viper"

type config struct {
	dbPath string
}

var c *config

func Init() {
	viperInit()
	c = &config{}

	c.dbPath = viper.GetString("DB_PATH")
}

func viperInit() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func GetDBPath() string {
	return c.dbPath
}
