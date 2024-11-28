package main

import (
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/payment/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
)

func main() {
	conf.ViperInit()
	dao.MysqlInit()
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Println("payment service register err :", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8086")

	svc := paymentservice.NewServer(new(PaymentServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(
			&registry.Info{
				ServiceName: "payment",
				Weight:      1,
			},
		),
		server.WithServiceAddr(addr),
	)

	err = svc.Run()
	if err != nil {
		log.Println("payment service run err : ", err)
	}
}
