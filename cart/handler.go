package main

import (
	"context"
	"errors"
	cart "github.com/Camelia-hu/gomall/cart/kitex_gen/cart"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/module"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log"
	"strconv"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	span := trace.SpanFromContext(ctx)
	var addCart module.Cart
	addCart.Uid = uint(req.UserId)
	addCart.ProductId = uint(req.Item.ProductId)
	addCart.Quantity = int(req.Item.Quantity)
	exist, _ := dao.Rdb.HExists(ctx, strconv.Itoa(int(req.UserId)), strconv.Itoa(int(req.Item.ProductId))).Result()
	if exist {
		dao.Rdb.HIncrBy(ctx, strconv.Itoa(int(req.UserId)), strconv.Itoa(int(req.Item.ProductId)), int64(req.Item.Quantity))
	} else {
		dao.Rdb.HSet(ctx, strconv.Itoa(int(req.UserId)), strconv.Itoa(int(req.Item.ProductId)), req.Item.Quantity)
	}
	err = dao.DB.Create(&addCart).Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		log.Println("add item err : ", err)
		return nil, err
	}
	resp = &cart.AddItemResp{}
	return resp, nil
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	span := trace.SpanFromContext(ctx)
	var carts []module.Cart
	result, err := dao.Rdb.HGetAll(ctx, strconv.Itoa(int(req.UserId))).Result()
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		err = dao.DB.Where("uid = ?", req.UserId).Find(&carts).Error
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
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
	resp = &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
			Items:  []*cart.CartItem{},
		},
	}
	for key, value := range result {
		var oneItem *cart.CartItem
		ikey, _ := strconv.Atoi(key)
		oneItem.ProductId = uint32(ikey)
		ivalue, _ := strconv.Atoi(value)
		oneItem.Quantity = int32(ivalue)
		resp.Cart.Items = append(resp.Cart.Items, oneItem)
	}
	return resp, nil
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	span := trace.SpanFromContext(ctx)
	uid := strconv.Itoa(int(req.UserId))
	result, err := dao.Rdb.Del(ctx, uid).Result()
	if err != nil || result <= 0 {
		span.SetStatus(codes.Error, "redis delete err")
		return nil, errors.New("redis delete err")
	}
	err = dao.DB.Where("uid = ?", req.UserId).Delete(&module.Cart{}).Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	resp = &cart.EmptyCartResp{}
	return resp, nil
}
