package main

import (
	"context"
	"errors"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/module"
	payment "github.com/Camelia-hu/gomall/payment/kitex_gen/payment"
	"gorm.io/gorm"
	"strconv"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	var oldCredit module.CreditCard

	err = dao.DB.Where("uid = ? and credit_card_number = ?", req.UserId, req.CreditCard.CreditCardNumber).First(&oldCredit).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("银行卡号输入错误")
	}
	if oldCredit.CreditCardCvv != req.CreditCard.CreditCardCvv {
		return nil, errors.New("校验码错误")
	}
	if err != nil {
		return nil, err
	}
	oldCredit.Money = oldCredit.Money - req.Amount
	if oldCredit.Money < 0 {
		return nil, errors.New("支付失败，余额不足")
	}
	Payment := module.Payment{
		Model:            gorm.Model{},
		CreditCardNumber: req.CreditCard.CreditCardNumber,
		Amount:           req.Amount,
		OrderId:          req.OrderId,
		UserId:           req.UserId,
	}
	err = dao.DB.Create(&Payment).Error
	err = dao.DB.Delete(&module.Order{}, req.OrderId).Error
	err = dao.DB.Model(&oldCredit).Update("money", oldCredit.Money).Error
	if err != nil {
		return nil, err
	}
	resp = &payment.ChargeResp{TransactionId: strconv.Itoa(int(Payment.ID))}
	return resp, nil
}

// CreateCredit implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CreateCredit(ctx context.Context, req *payment.CreateCreditReq) (resp *payment.CreateCreditResp, err error) {
	err = dao.DB.Create(&module.CreditCard{
		Uid:                       uint32(req.Uid),
		Money:                     req.Money,
		CreditCardNumber:          req.CreditCardNumber,
		CreditCardCvv:             req.CreditCardCvv,
		CreditCardExpirationYear:  req.CreditCardExpirationYear,
		CreditCardExpirationMonth: req.CreditCardExpirationMonth,
	}).Error
	if err != nil {
		return nil, err
	}
	resp = &payment.CreateCreditResp{Is: true}
	return resp, nil
}