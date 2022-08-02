package usecase

import (
	"github.com/Saucon/simple-bank/account/internal/model"
	"github.com/Saucon/simple-bank/account/pkg/log"
)

type IAccountUsecase interface {
	CreateAccount(account model.Accounts) error
	UpdateAccount()
	DeactivateAccount()
}

type accountUsecase struct {
	log *log.LogCustom
}

func (a accountUsecase) CreateAccount(account model.Accounts) error {
	//TODO implement me
	return nil
}

func (a accountUsecase) UpdateAccount() {
	//TODO implement me
	panic("implement me")
}

func (a accountUsecase) DeactivateAccount() {
	//TODO implement me
	panic("implement me")
}

func NewAccountUsecase(a *log.LogCustom) IAccountUsecase {
	return accountUsecase{
		log: a,
	}
}
