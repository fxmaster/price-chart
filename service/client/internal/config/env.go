package config

import (
	"price-chart/pkg/util"
	"sync"
)

var (
	env  *environment
	err  error
	once sync.Once
)

type environment struct {
	App struct {
		Port int `env:"APP_PORT"`
	}
	MongoParams struct {
		Host       string `env:"MONGO_HOST"`
		DB         string `env:"MONGO_DB"`
		Collection string `env:"MONGO_COLLECTION"`
		User       string `env:"MONGO_USER"`
		Password   string `env:"MONGO_PASSWORD"`
		Port       int    `env:"MONGO_PORT"`
	}
}

func LoadEnv(path string) error {
	once.Do(func() {
		env = &environment{}
		err = util.LoadEnv(env, path)
	})

	return err
}

func Environment() *environment {
	return env
}
