package main

import (
	"github.com/soheilsleep/golang-clean-web-api/api"
	"github.com/soheilsleep/golang-clean-web-api/config"
	"github.com/soheilsleep/golang-clean-web-api/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedisClient()
	api.InitServer(cfg)
}
