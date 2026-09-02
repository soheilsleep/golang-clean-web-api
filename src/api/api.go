package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/soheilsleep/golang-clean-web-api/api/routers"
	"github.com/soheilsleep/golang-clean-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	api := r.Group("/api")
	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
	}
	//err := r.Run(":5005")
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))

	if err != nil {
		panic(err)
	}

}
