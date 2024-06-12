package model

import (
	"github.com/shopspring/decimal"
)

type Account struct {
	BaseEntity
	Name          string          `json:"name" validate:"required"`
	UserId        int64           `json:"userId" validate:"required"`
	ExchangeId    int64           `json:"exchangeId" validate:"required"`
	Capital       decimal.Decimal `json:"capital" validate:"required"`
	LockedCapital decimal.Decimal `json:"lockedCapital"`
}
