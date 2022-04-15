package app

import (
	"log"
	"path/filepath"
	"price-chart/service/parser/internal/config"
)

func InitEnv(path string) {
	err := config.LoadEnv(filepath.FromSlash(path))
	if err != nil {
		log.Fatalln(err)
	}
}
