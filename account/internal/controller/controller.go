package controller

import (
	"github.com/Saucon/simple-bank/account/internal/model"
	"github.com/Saucon/simple-bank/account/internal/usecase"
	"github.com/Saucon/simple-bank/account/pkg/log"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	log *log.LogCustom

	accntUsecase usecase.IAccountUsecase
}

func NewAccountHandler(a *log.LogCustom, b usecase.IAccountUsecase) AccountHandler {
	return AccountHandler{
		log:          a,
		accntUsecase: b,
	}
}

func (ah *AccountHandler) CreateAccount(c *gin.Context) {
	var request model.RequestCreateAccount

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = ah.accntUsecase.CreateAccount(request.Accounts)

	c.JSON(201, model.ResponseCreateAccount{
		ResponseCode:    "201XX00",
		ResponseMessage: "Successful",
		Body:            request.Accounts,
	})
	return
}

func (ah *AccountHandler) UpdateAccount(c *gin.Context) {
	var request model.RequestUpdateAccount

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(500, err)
		return
	}

	ah.log.Success(request, nil, nil, nil, "controller - UpdateAccount", "", nil)
	c.JSON(201, model.ResponseCreateAccount{
		ResponseCode:    "200XX00",
		ResponseMessage: "Successful",
		Body:            request.Accounts,
	})
	return
}

func (ah *AccountHandler) DeactivateAccount(c *gin.Context) {
	var request model.RequestDeactivateAccount

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(201, model.ResponseDeactivateAccount{
		ResponseCode:    "200XX00",
		ResponseMessage: "Successful",
	})
	return
}
