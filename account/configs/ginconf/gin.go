package ginconf

import (
	"github.com/Saucon/simple-bank/account/pkg/env"
	"github.com/Saucon/simple-bank/account/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type routerConfig struct {
	Gin *gin.Engine
}

func NewRouter() *routerConfig {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"responseCode": "40404", "responseMessage": "Invalid Path"})
	})

	return &routerConfig{
		Gin: router,
	}
}

func (rc *routerConfig) GroupingRouter(logger *log.LogCustom, env env.ServerConfig, handler ...func(c *gin.Context)) *gin.Engine {
	if len(handler) == 0 {
		logger.Fatal(errors.New(""), "handler is null", nil)
	}

	api := rc.Gin.Group("/" + env.Version + "/simple-bank/internal")
	api.POST("/account", handler[0])

	return rc.Gin
}
