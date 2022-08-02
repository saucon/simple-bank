package controller

import (
	"github.com/Saucon/simple-bank/account/internal/model"
	"github.com/Saucon/simple-bank/account/pkg/log"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	log *log.LogCustom
}

func NewAccountHandler(a *log.LogCustom) AccountHandler {
	return AccountHandler{
		log: a,
	}
}

func (ah *AccountHandler) CreateAccount(c *gin.Context) {
	var request model.RequestCreateAccount

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(201, model.ResponseCreateAccount{
		ResponseCode:    "201XX00",
		ResponseMessage: "Successful",
		Body:            request.Accounts,
	})
	return
}
