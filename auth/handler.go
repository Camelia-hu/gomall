package main

import (
	"context"
	"errors"
	"github.com/Camelia-hu/gomall/auth/kitex_gen/auth"
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/module"
	"github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log"
	"time"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	span := trace.SpanFromContext(ctx)
	if !span.SpanContext().IsValid() {
		return nil, errors.New("span invalid")
	}
	myAccessClaims := module.MyClaims{
		Id: int(req.UserId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * 60 * time.Minute)),
		},
	}
	myRefreshClaims := module.MyClaims{
		Id: int(req.UserId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
		},
	}
	resp = &auth.DeliveryResp{}
	resp.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, myAccessClaims).SignedString([]byte(conf.Conf.GetString("data.jwt.key")))
	resp.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, myRefreshClaims).SignedString([]byte(conf.Conf.GetString("data.jwt.key")))
	if err != nil {
		span.SetStatus(codes.Error, "token deliver err")
		return nil, err
	}
	return resp, nil
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	span := trace.SpanFromContext(ctx)
	if !span.SpanContext().IsValid() {
		return nil, errors.New("span invalid")
	}
	myClaims := new(module.MyClaims)
	token, err := jwt.ParseWithClaims(req.Token, myClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf.GetString("data.jwt.key")), nil
	})
	if err != nil {
		span.SetStatus(codes.Error, "token parse err")
		log.Println("1 ", err)
		return nil, err
	}
	if !token.Valid {
		span.SetStatus(codes.Error, "token expired")
		return nil, errors.New("token 过期喵～")
	}
	resp = &auth.VerifyResp{}
	resp.Res = true

	return resp, nil
}

// ReFreshTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ReFreshTokenByRPC(ctx context.Context, req *auth.RefreshReq) (resp *auth.RefreshResp, err error) {
	span := trace.SpanFromContext(ctx)
	var myClaims module.MyClaims
	refreshToken, err := jwt.ParseWithClaims(req.RefreshToken, &myClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf.GetString("data.jwt.key")), nil
	})
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	if !refreshToken.Valid {
		span.SetStatus(codes.Error, "refreshToken 过期 too 喵～")
		return nil, errors.New("refreshToken 过期 too 喵～")
	}
	newrefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, module.MyClaims{
		Id: myClaims.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
		},
	}).SignedString([]byte(conf.Conf.GetString("data.jwt.key")))
	newaccessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, module.MyClaims{
		Id: myClaims.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	}).SignedString([]byte(conf.Conf.GetString("data.jwt.key")))
	resp = &auth.RefreshResp{
		AccessToken:  newaccessToken,
		RefreshToken: newrefreshToken,
	}
	return resp, nil
}
