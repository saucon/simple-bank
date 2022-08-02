package usecase

import "github.com/Saucon/simple-bank/account/pkg/log"

type IAccountUsecase interface {
	CreateAccount()
	UpdateAccount()
	DeactivateAccount()
}

type accountUsecase struct {
	log *log.LogCustom
}

func (a accountUsecase) CreateAccount() {
	//TODO implement me
	panic("implement me")
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
