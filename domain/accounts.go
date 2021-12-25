package domain

import (
	"github.com/leslesnoa/go-microservices-demo/dto"
	"github.com/leslesnoa/go-microservices-demo/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
}
