package main

import (
	"context"
	"errors"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/module"
	product "github.com/Camelia-hu/gomall/product/kitex_gen/product"
	"gorm.io/gorm"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// TODO: Your code here...
	return
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	var Product module.Product
	err = dao.DB.Where("id = ?", req.Id).First(&Product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("未找到该商品信息喵～")
	}
	if err != nil {
		return nil, err
	}
	return
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	var Products []*product.Product
	err = dao.DB.Where("name like ?", "%"+req.Query+"%").Find(Products).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("没有记录喵～")
	}
	if err != nil {
		return nil, err
	}
	resp.Results = Products
	return resp, nil
}

// CreateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) CreateProduct(ctx context.Context, req *product.CreateReq) (resp *product.CreateResp, err error) {
	if req.Name == "" || req.Categories == nil || &req.Price == nil {
		return nil, errors.New("请输入商品名称，价格以或分类标签喵～")
	}
	var Product product.Product
	var newProduct product.Product
	Product.Name = req.Name
	Product.Price = req.Price
	Product.Categories = req.Categories
	Product.Picture = req.Picture
	Product.Description = req.Description
	dao.DB.Create(&Product)
	err = dao.DB.Where("name = ?", Product.Name).First(&newProduct).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("数据库插入失败喵～")
	}
	resp.Id = newProduct.Id
	return resp, nil
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteReq) (resp *product.DeleteResp, err error) {
	err = dao.DB.Delete(&product.Product{}, req.Id).Error
	if err != nil {
		resp.Is = false
		return resp, err
	}
	resp.Is = true
	return resp, nil
}
