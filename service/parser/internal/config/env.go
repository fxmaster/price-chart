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
