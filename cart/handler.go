package main

import (
	"context"
	cart "github.com/Camelia-hu/gomall/cart/kitex_gen/cart"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/module"
	"log"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	var addCart module.Cart
	addCart.Uid = uint(req.UserId)
	addCart.ProductId = uint(req.Item.ProductId)
	addCart.Quantity = int(req.Item.Quantity)
	err = dao.DB.Create(&addCart).Error
	if err != nil {
		log.Println("add item err : ", err)
		return nil, err
	}
	resp = &cart.AddItemResp{}
	return resp, nil
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	var carts []module.Cart
	err = dao.DB.Where("uid = ?", req.UserId).Find(&carts).Error
	if err != nil {
		log.Println("get cart err : ", err)
		return nil, err
	}
	resp = &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
			Items:  []*cart.CartItem{},
		},
	}
	for _, oneCart := range carts {
		item := &cart.CartItem{
			ProductId: uint32(oneCart.ProductId),
			Quantity:  int32(oneCart.Quantity),
		}
		resp.Cart.Items = append(resp.Cart.Items, item)
	}
	return resp, nil
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	err = dao.DB.Where("uid = ?", req.UserId).Delete(&module.Cart{}).Error
	if err != nil {
		return nil, err
	}
	resp = &cart.EmptyCartResp{}
	return resp, nil
}
