package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig() (*conf, error) {
	envFile := filepath.Join(".", ".env")

	viper.SetConfigType("env")
	viper.SetConfigFile(envFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	checkRequiredEnvs()

	return cfg, err
}

func checkRequiredEnvs() {
	requiredEnvs := []string{
		"SERVER_PORT",
		"DB_DRIVER",
		"DB_HOST",
		"DB_USER",
		"DB_PASS",
		"DB_NAME",
		"DB_PORT",
		"DB_URL",
		"WEB_URL",
		"JWT_SECRET",
	}

	for _, env := range requiredEnvs {
		if viper.Get(env) == nil {
			panic(env + " env not set")
		}
	}
}
