package model

import (
	"github.com/shopspring/decimal"
)

type Account struct {
	BaseEntity
	Name          string          `json:"name" validate:"required"`
	UserId        int             `json:"user_id" validate:"required"`
	ExchangeId    int             `json:"exchange_id" validate:"required"`
	Capital       decimal.Decimal `json:"capital" validate:"required"`
	LockedCapital decimal.Decimal `json:"locked_capital"`
}
