package main

import (
	"github.com/Saucon/simple-bank/account/configs/ginconf"
	"github.com/Saucon/simple-bank/account/pkg/env"
	"log"
)

func main() {
	// get env
	env.NewEnv("account/.env")
	cfg := env.Config

	router := ginconf.NewRouter()

	if err := router.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Fatalf("Something was wrong with "+cfg.Host+":"+cfg.Port, err)
	}
}
