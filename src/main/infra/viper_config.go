package infra

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseStruct struct {
	Host   string `mapstructure:"DB_HOST"`
	Port   string `mapstructure:"DB_PORT"`
	User   string `mapstructure:"DB_USER"`
	Pass   string `mapstructure:"DB_PASS"`
	Dbname string `mapstructure:"DB_NAME"`
}

func LoadEnv(path string) (db DatabaseStruct, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("error Loading .env file")
		log.Fatal(err.Error())
	}

	err = viper.Unmarshal(&db)
	if err != nil {
		log.Fatal("error Unmarshal .env file")
		log.Fatal(err.Error())
	}

	return
}
