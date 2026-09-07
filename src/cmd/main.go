package main

import (
	"log"

	"github.com/soheilsleep/golang-clean-web-api/api"
	"github.com/soheilsleep/golang-clean-web-api/config"
	"github.com/soheilsleep/golang-clean-web-api/data/cache"
	"github.com/soheilsleep/golang-clean-web-api/data/db"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedisClient()

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
