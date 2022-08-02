package main

import (
	"github.com/Saucon/simple-bank/account/configs/ginconf"
	"github.com/Saucon/simple-bank/account/internal/controller"
	"github.com/Saucon/simple-bank/account/internal/usecase"
	"github.com/Saucon/simple-bank/account/pkg/env"
	"github.com/Saucon/simple-bank/account/pkg/log"
)

func main() {
	// get env
	env.NewEnv("account/.env")
	cfg := env.Config

	logger := log.NewLogCustom(cfg)

	router := ginconf.NewRouter()

	accountUsecase := usecase.NewAccountUsecase(logger)
	accountHandler := controller.NewAccountHandler(logger, accountUsecase)

	router.Gin = router.GroupingRouter(logger, cfg, accountHandler.CreateAccount)
	if err := router.Gin.Run(cfg.Host + ":" + cfg.Port); err != nil {
		logger.Fatal(err, "Something was wrong with "+cfg.Host+":"+cfg.Port, nil)
	}
}
