package app

import (
	"context"
	"log"
	"path/filepath"
	"price-chart/service/client/internal/config"
	"price-chart/service/client/internal/model/migration"
	"time"
)

func InitEnv(path string) {
	err := config.LoadEnv(filepath.FromSlash(path))
	if err != nil {
		log.Fatalln(err)
	}
}

func InitDB() {
	log.Println("init db...")
	_, err := config.MongoClient()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connected to db")
}

func Migrate() {
	log.Println("started migration...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	client, err := config.MongoClient()
	if err != nil {
		log.Fatalln(err)
	}

	err = migration.CreateCollection(ctx, client, config.Environment().MongoParams.Collection)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("migration finished")
}
