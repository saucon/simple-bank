package main

import (
	"github.com/Saucon/simple-bank/account/pkg/env"
	"github.com/Saucon/simple-bank/customer/configs/ginconf"
	"log"
)

func main() {
	// get env
	env.NewEnv("customer/.env")
	cfg := env.Config

	router := ginconf.NewRouter()

	if err := router.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Fatalf("Something was wrong with "+cfg.Host+":"+cfg.Port, err)
	}
}
