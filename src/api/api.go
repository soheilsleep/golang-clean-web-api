package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/soheilsleep/golang-clean-web-api/api/middlewares"
	"github.com/soheilsleep/golang-clean-web-api/api/routers"
	"github.com/soheilsleep/golang-clean-web-api/api/validations"
	"github.com/soheilsleep/golang-clean-web-api/config"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	RegisterValidators()
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Recovery(), gin.Logger(), middlewares.LimitByRequest() /*middlewares.TestMiddleware()*/)
	RegisterRoutes(r)
	//err := r.Run(":5005")
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))

	if err != nil {
		panic(err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validations.PasswordValidator, true)
	}
}
