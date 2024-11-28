package main

import (
	"context"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/module"
	order "github.com/Camelia-hu/gomall/order/kitex_gen/order"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	var ReqOrder module.Order
	req.Address = &order.Address{}
	address := req.Address.Country + req.Address.State + req.Address.City + req.Address.StreetAddress
	ReqOrder = module.Order{
		Model:        gorm.Model{},
		Uid:          req.UserId,
		UserCurrency: req.UserCurrency,
		Address:      address,
		Email:        req.Email,
	}
	err = dao.DB.Create(&ReqOrder).Error
	if err != nil {
		return nil, err
	}
	for _, item := range req.OrderItems {
		oneItem := module.OrderItem{
			Model:     gorm.Model{},
			OrderID:   ReqOrder.ID,
			Cost:      item.Cost,
			ProductId: item.Item.ProductId,
			Quantity:  item.Item.Quantity,
		}
		err = dao.DB.Create(&oneItem).Error
		if err != nil {
			log.Println("order item create err : ", err)
			return nil, err
		}
	}
	resp = &order.PlaceOrderResp{Order: &order.OrderResult{OrderId: strconv.Itoa(int(ReqOrder.ID))}}
	return resp, nil
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	var orders []*module.Order
	err = dao.DB.Where("uid = ?", req.UserId).Find(&orders).Error
	if err != nil {
		log.Println("find one_order list err : ", err)
		return nil, err
	}
	var Orders []*order.Order
	for _, o := range orders {
		oneOrder := &order.Order{
			OrderId:      strconv.Itoa(int(o.ID)),
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

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// TODO: Your code here...
	return
}
