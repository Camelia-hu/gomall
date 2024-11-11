package main

import (
	"context"
	"errors"
	"github.com/Camelia-hu/gomall/auth/kitex_gen/auth"
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/module"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	myAccessClaims := module.MyClaims{
		Id: int(req.UserId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}
	myRefreshClaims := module.MyClaims{
		Id: int(req.UserId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
		},
	}
	resp.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, myAccessClaims).SignedString(conf.Conf.GetString("data.jwt.key"))
	resp.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, myRefreshClaims).SignedString(conf.Conf.GetString("data.jwt.key"))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	myClaims := new(module.MyClaims)
	token, err := jwt.ParseWithClaims(req.Token, myClaims, func(token *jwt.Token) (interface{}, error) {
		return conf.Conf.GetString("data.jwt.key"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token 过期喵～")
	}
	resp.Res = true
	return resp, nil
}

// ReFreshTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ReFreshTokenByRPC(ctx context.Context, req *auth.RefreshReq) (resp *auth.RefreshResp, err error) {
	var myClaims module.MyClaims
	refreshToken, err := jwt.ParseWithClaims(req.RefreshToken, &myClaims, func(token *jwt.Token) (interface{}, error) {
		return conf.Conf.GetString("data.jwt.key"), nil
	})
	if err != nil {
		return nil, err
	}
	if !refreshToken.Valid {
		return nil, errors.New("refreshToken 过期 too 喵～")
	}
	accessToken, err := jwt.ParseWithClaims(req.AccessToken, &myClaims, func(token *jwt.Token) (interface{}, error) {
		return conf.Conf.GetString("data.jwt.key"), nil
	})
	if err != nil {
		return nil, errors.New("为啥accessToken会有问题？？？")
	}
	resp.RefreshToken, err = refreshToken.SignedString(conf.Conf.GetString("data.jwt.key"))
	resp.AccessToken, err = accessToken.SignedString(conf.Conf.GetString("data.jwt.key"))
	return resp, nil
}
