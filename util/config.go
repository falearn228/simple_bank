package util

import "github.com/spf13/viper"

// Конфиг содержит все конфигурации приложения. Значения
// читаются с помощью Viper'a из конфиг-файла или env variables
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	Address  string `mapstructure:"ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
