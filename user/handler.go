package main

import (
	"context"
	"errors"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/module"
	user "github.com/Camelia-hu/gomall/user/kitex_gen/user"
	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("请输入邮箱地址或密码喵~")
	}
	var usr module.User
	err1 := dao.DB.Where("email = ?", req.Email).First(&usr).Error
	if !errors.Is(err1, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名已存在喵~")
	}
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("两次密码输入不一致喵~")
	}
	usr.Email = req.Email
	usr.Password = req.Password
	dao.DB.Create(&usr)
	var newusr module.User
	dao.DB.Where("email = ?", usr.Email).First(&newusr)
	resp.UserId = int32(newusr.ID)
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	var usr module.User
	if req.Password == "" || req.Email == "" {
		return nil, errors.New("请输入邮箱地址或密码喵～")
	}
	err = dao.DB.Where("email = ?", req.Email).First(&usr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("该用户名不存在喵～")
	}
	if req.Password != usr.Password {
		return nil, errors.New("密码输入错误喵～")
	}
	resp.UserId = int32(usr.ID)
	return resp, nil
}
