package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/soheilsleep/golang-clean-web-api/api/routers"
	"github.com/soheilsleep/golang-clean-web-api/api/validations"
	"github.com/soheilsleep/golang-clean-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			err.Error()
		}
	}
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
