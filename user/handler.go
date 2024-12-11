package main

import (
	"context"
	"errors"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/module"
	user "github.com/Camelia-hu/gomall/user/kitex_gen/user"
	"github.com/Camelia-hu/gomall/utils"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

//如果后续用户表里面出现个人信息类字段，应该将其储存在redis里面，目前的字段没有必要存储在redis里面

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	span := trace.SpanFromContext(ctx)
	if req.Email == "" || req.Password == "" {
		span.SetStatus(codes.Error, "请输入邮箱地址或密码喵~")
		return nil, errors.New("请输入邮箱地址或密码喵~")
	}
	var usr module.User
	err1 := dao.DB.Where("email = ?", req.Email).First(&usr).Error
	if !errors.Is(err1, gorm.ErrRecordNotFound) {
		span.SetStatus(codes.Error, "用户名已存在喵")
		return nil, errors.New("用户名已存在喵~")
	}
	if req.Password != req.ConfirmPassword {
		span.SetStatus(codes.Error, "两次密码输入不一致喵")
		return nil, errors.New("两次密码输入不一致喵~")
	}
	usr.Email = req.Email
	usr.Password = req.Password
	salt := utils.GenerateSalt()
	usr.Salt = salt
	usr.Password = utils.HashPassword(usr.Password, salt)
	dao.DB.Create(&usr)
	var newusr module.User
	dao.DB.Where("email = ?", usr.Email).First(&newusr)
	resp = &user.RegisterResp{}
	resp.UserId = int32(newusr.ID)
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	span := trace.SpanFromContext(ctx)
	var usr module.User
	if req.Password == "" || req.Email == "" {
		span.SetStatus(codes.Error, "请输入邮箱地址或密码喵")
		return nil, errors.New("请输入邮箱地址或密码喵～")
	}
	err = dao.DB.Where("email = ?", req.Email).First(&usr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		span.SetStatus(codes.Error, "该用户名不存在喵")
		return nil, errors.New("该用户名不存在喵～")
	}
	if utils.HashPassword(req.Password, usr.Salt) != usr.Password {
		span.SetStatus(codes.Error, "密码输入错误喵～")
		return nil, errors.New("密码输入错误喵～")
	}
	resp = &user.LoginResp{}
	resp.UserId = int32(usr.ID)
	return resp, nil
}
