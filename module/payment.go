package module

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CreditCardNumber string  `json:"creditCardNumber"`
	Amount           float32 `json:"amount"`
	OrderId          string  `json:"orderId"`
	UserId           uint32  `json:"userId"`
}

type CreditCard struct {
	gorm.Model
	Uid                       uint32  `json:"uid"`
	Money                     float32 `json:"money"`
	CreditCardNumber          string  `json:"creditCardNumber"`
	CreditCardCvv             int32   `json:"creditCardCvv"`
	CreditCardExpirationYear  int32   `json:"creditCardExpirationYear"`
	CreditCardExpirationMonth int32   `json:"creditCardExpirationMonth"`
}
