package main

import (
	"context"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/module"
	order "github.com/Camelia-hu/gomall/order/kitex_gen/order"
	"log"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	var orders []*module.Order
	address := req.Address.Country + req.Address.State + req.Address.City + req.Address.StreetAddress
	for _, item := range req.OrderItems {
		Order := &module.Order{
			Uid:          req.UserId,
			UserCurrency: req.UserCurrency,
			Address:      address,
			Email:        req.Email,
		}
		orders = append(orders, Order)
	}
	err = dao.DB.Create(orders).Error
	if err != nil {
		log.Println("order create err : ", err)
		return nil, err
	}
	resp = &order.PlaceOrderResp{Order: &order.OrderResult{OrderId: Uuid}}
	return resp, nil
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	var orders []module.Order
	err = dao.DB.Where("uid = ?", req.UserId).Find(&orders).Error
	if err != nil {
		log.Println("find one_order list err : ", err)
		return nil, err
	}
	var Orders []*order.Order
	for _, o := range orders {
		oneOrder := &order.Order{
			OrderId:      o.UUid,
			UserId:       o.Uid,
			UserCurrency: o.UserCurrency,
			Email:        o.Email,
		}
		//here
		Orders = append(Orders, oneOrder)
	}
	resp = &order.ListOrderResp{}
	resp.Orders = Orders
	return resp, nil
}

// MarkOrderPaid implements the OrderServiceImpl interface.
