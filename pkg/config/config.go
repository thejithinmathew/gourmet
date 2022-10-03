package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	User           string `json:"user"`
	Pass           string `json:"pass"`
	URI            string `json:"mongo_uri"`
	DBName         string `json:"db_name"`
	CollectionName string `json:"col_name"`
}

var c = Config{}

func Init() {
	var configData []byte
	configData, err := os.ReadFile("C:/Users/theji/go/src/github.com/thejithinmathew/gourmet/config.json")
	if err != nil {
		panic("cannot load config, panic!")
	}
	mErr := json.Unmarshal(configData, &c)
	if mErr != nil {
		panic("cannot load config, panic!")
	}
}

func Get() Config {
	return c
}
