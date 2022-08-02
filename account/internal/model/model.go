package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	SavingsAccount string = "SAVINGS_ACCOUNT"
	CurrentAccount        = "CURRENT_ACCOUNT"
)

type Accounts struct {
	Id             uint           `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt      time.Time      `json:"createdAt,omitempty"`
	UpdatedAt      time.Time      `json:"updatedAt,omitempty"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	BalanceAmount  BalanceAmount  `json:"balanceAmount" gorm:"type:JSONB NULL DEFAULT '{}'::JSONB" validate:"required,len=8" `
	CustomerId     string         `json:"customerId"`
	Type           string         `json:"type"`
	AccountNumber  string         `json:"accountNumber" gorm:"index"`
	AccountStatus  string         `json:"accountStatus"`
	CreationStatus string         `json:"creationStatus"`
}

type BalanceAmount struct {
	Value        string `json:"value"`
	Currency     string `json:"currency"`
	BalanceFloat float64
}

type RequestCreateAccount struct {
	Accounts
}

type ResponseCreateAccount struct {
	ResponseCode    string   `json:"responseCode"`
	ResponseMessage string   `json:"responseMessage"`
	Body            Accounts `json:"body"`
}

type RequestUpdateAccount struct {
	Accounts
}

type ResponseUpdateAccount struct {
	ResponseCode    string   `json:"responseCode"`
	ResponseMessage string   `json:"responseMessage"`
	Body            Accounts `json:"body"`
}

type RequestGetAccount struct {
	AccountNumber string `json:"accountNumber"`
}

type ResponseGetAccount struct {
	ResponseCode    string   `json:"responseCode"`
	ResponseMessage string   `json:"responseMessage"`
	Body            Accounts `json:"body"`
}

type RequestDeactivateAccount struct {
	AccountNumber string `json:"accountNumber"`
}

type ResponseDeactivateAccount struct {
	ResponseCode    string   `json:"responseCode"`
	ResponseMessage string   `json:"responseMessage"`
	Body            Accounts `json:"body"`
}
