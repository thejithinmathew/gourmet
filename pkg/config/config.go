package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	User string
	Pass string
	URI  string
}

var c = Config{}

func Init() {
	envconfig.MustProcess("gourmet", &c)
}

func Get() Config {
	return c
}
